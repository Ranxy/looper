package bind

import (
	"container/list"
	"fmt"

	"github.com/Ranxy/looper/buildin"
	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

type Binder struct {
	Diagnostics *diagnostic.DiagnosticBag
	scope       *BoundScope
	function    *symbol.FunctionSymbol

	_labelCount  int
	_looperStack list.List //stack[breakContinue]
}

type breakContinue struct {
	BreakLabel    *BoundLabel
	ContinueLabel *BoundLabel
}

func NewBinder(parent *BoundScope, function *symbol.FunctionSymbol) *Binder {
	res := &Binder{
		Diagnostics: diagnostic.NewDiagnostics(),
		scope:       NewBoundScope(parent),
		function:    function,
	}
	if function != nil {
		for _, p := range function.Parameter {
			res.scope.TryDeclareVariable(p)
		}
	}

	return res
}

func BindErrorStatement() Boundstatement {
	return NewBoundExpressStatements(NewBoundErrorExpression())
}

func BindGlobalScope(previous *BoundGlobalScope, s *syntax.CompliationUnit) *BoundGlobalScope {
	parentScope := CreateParentScope(previous)
	binder := NewBinder(parentScope, nil)

	for _, member := range s.Statements {
		if function, ok := member.(*syntax.FunctionDeclarationSyntax); ok {
			binder.BindFunctionDeclaration(function)
		}
	}
	statements := make([]Boundstatement, 0)

	for _, member := range s.Statements {
		if gs, ok := member.(*syntax.GlobalStatement); ok {
			bs := binder.BindStatement(gs.Statement)
			statements = append(statements, bs)
		}
	}
	functions := binder.scope.GetDeclareFunctions()
	variables := binder.scope.GetDeclareVariables()
	diagnostics := binder.Diagnostics

	if previous != nil {
		diagnostics = diagnostics.Merge(previous.Diagnostic)
	}

	return NewBoundGlobalScope(previous, diagnostics, functions, variables, statements)
}

func CreateParentScope(previous *BoundGlobalScope) *BoundScope {

	stack := list.New()

	for previous != nil {
		stack.PushBack(previous)
		previous = previous.Previous
	}

	parent := CreateRootScope()

	for stack.Len() > 0 {
		pvi := stack.Back()
		if pvi != nil {
			stack.Remove(pvi)
		}
		pv := pvi.Value.(*BoundGlobalScope)
		scope := NewBoundScope(parent)
		for _, f := range pv.Functions {
			scope.TryDeclareFunction(f)
		}
		for _, v := range pv.Variables {
			scope.TryDeclareVariable(v)
		}
		parent = scope
	}
	return parent
}

func CreateRootScope() *BoundScope {
	res := NewBoundScope(nil)
	for _, f := range buildin.AllBuildinFunc {
		_ = res.TryDeclareFunction(f)
	}
	return res
}

func (b *Binder) BindExpressionAndCheckType(express syntax.Express, targetType *symbol.TypeSymbol) BoundExpression {
	result := b.BindExpressionInternal(express)
	if result.Type() != symbol.TypeError && targetType != symbol.TypeError && result.Type() != targetType {
		b.Diagnostics.CannotConvert(syntax.SyntaxNodeSpan(express), result.Type(), targetType)
		return NewBoundErrorExpression()
	}
	return result
}

func (b *Binder) BindExpression(express syntax.Express, canBeUnit bool) BoundExpression {
	result := b.BindExpressionInternal(express)
	if !canBeUnit && result.Type() == symbol.TypeUnit {
		b.Diagnostics.ExpressionMustReturnValue(syntax.SyntaxNodeSpan(express))
		return NewBoundErrorExpression()
	}
	return result
}

func (b *Binder) BindExpressionInternal(express syntax.Express) BoundExpression {
	switch express.Kind() {
	case syntax.SyntaxKindLiteralExpress:
		return b.BindLiteralExpress(express.(*syntax.LiteralExpress))
	case syntax.SyntaxKindNameExpress:
		return b.BindNameExpress(express.(*syntax.NameExpress))
	case syntax.SyntaxKindUnaryExpress:
		return b.BindUnaryExpress(express.(*syntax.UnaryExpress))
	case syntax.SyntaxKindBinaryExpress:
		return b.BindBinaryOperator(express.(*syntax.BinaryExpress))
	case syntax.SyntaxKindParenthesizedExpress:
		return b.BindExpressionInternal(express.(*syntax.ParenthesisExpress).Expr)
	case syntax.SyntaxKindAssignmentExpress:
		return b.BindAssignmentExpress(express.(*syntax.AssignmentExpress))
	case syntax.SyntaxKindCallExpress:
		return b.BindCallExpression(express.(*syntax.CallExpress))
	case syntax.SyntaxKindUnitExpress:
		return b.BindUnitExpression(express.(*syntax.UnitExpress))
	default:
		panic(fmt.Sprintf("unexceped expresss %q", express.Kind()))
	}
}

func (b *Binder) BindStatement(s syntax.Statement) Boundstatement {
	switch s.Kind() {
	case syntax.SyntaxKindBlockStatement:
		return b.BindBlockStatement(s.(*syntax.BlockStatement))
	case syntax.SyntaxKindVariableDeclaration:
		return b.BindVariableDeclaration(s.(*syntax.VariableDeclarationSyntax))
	case syntax.SyntaxKindIfStatement:
		return b.BindIfStatement(s.(*syntax.IfStatement))
	case syntax.SyntaxKindWhileStatement:
		return b.BindWhileStatement(s.(*syntax.WhileStatement))
	case syntax.SyntaxkindForStatement:
		return b.BindForStatement(s.(*syntax.ForStatement))
	case syntax.SyntaxKindBreakStatement:
		return b.bindBreakStatement(s.(*syntax.BreakStatement))
	case syntax.SyntaxKindContinueStatement:
		return b.bindContinueStatement(s.(*syntax.ContinueStatement))
	case syntax.SyntaxKindReturnStatement:
		return b.BindReturnStatement(s.(*syntax.ReturnStatement))
	case syntax.SyntaxKindExpressStatement:
		return b.BindExpressionStatement(s.(*syntax.ExpressStatement))
	case syntax.SyntaxKindFunctionDeclaration:
		f := s.(*syntax.FunctionDeclarationSyntax)
		b.Diagnostics.FunctionNotTopLevel(f.Identifier.Span(), f.Identifier.Text)
		return NewBoundErrorExpression()
	default:
		panic(fmt.Sprintf("Unexceped syntax %s", s.Kind()))
	}
}

func (b *Binder) BindForStatement(s *syntax.ForStatement) Boundstatement {
	initCond := b.BindStatement(s.InitCondition)
	endCond := b.BindExpressionAndCheckType(s.EndCondition, symbol.TypeBool)
	updateCond := b.BindStatement(s.UpdateCondition)
	body, breakLabel, continueLabel := b.BindLoopBody(s.Body)
	return NewBoundForStatements(initCond, endCond, updateCond, body, breakLabel, continueLabel)
}

func (b *Binder) BindWhileStatement(s *syntax.WhileStatement) Boundstatement {
	condition := b.BindExpressionAndCheckType(s.Condition, symbol.TypeBool)
	body, breakLabel, continueLabel := b.BindLoopBody(s.Body)
	return NewBoundWhileStatements(condition, body, breakLabel, continueLabel)
}

func (b *Binder) BindLoopBody(body syntax.Statement) (res Boundstatement, breakLabel *BoundLabel, continueLabel *BoundLabel) {
	b._labelCount += 1
	breakLabel = NewBoundLabel(fmt.Sprintf("_break%d", b._labelCount))
	continueLabel = NewBoundLabel(fmt.Sprintf("_continue%d", b._labelCount))

	b._looperStack.PushFront(breakContinue{
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
	})
	boundBody := b.BindStatement(body)
	b._looperStack.Remove(b._looperStack.Front())

	return boundBody, breakLabel, continueLabel
}

func (b *Binder) bindBreakStatement(s *syntax.BreakStatement) Boundstatement {
	if b._looperStack.Len() == 0 {
		b.Diagnostics.ReportInvalidBreakOrContinue(s.Keywords.Span(), s.Keywords.Text)
		return BindErrorStatement()
	}

	breakLabel := b._looperStack.Front().Value.(breakContinue).BreakLabel

	return NewGotoSymbol(breakLabel)
}

func (b *Binder) bindContinueStatement(s *syntax.ContinueStatement) Boundstatement {
	if b._looperStack.Len() == 0 {
		b.Diagnostics.ReportInvalidBreakOrContinue(s.Keywords.Span(), s.Keywords.Text)
		return BindErrorStatement()
	}

	continueLabel := b._looperStack.Front().Value.(breakContinue).ContinueLabel

	return NewGotoSymbol(continueLabel)
}

func (b *Binder) BindReturnStatement(s *syntax.ReturnStatement) Boundstatement {

	expression := b.BindExpression(s.Express, true)

	if b.function == nil {
		b.Diagnostics.ReportInvalidReturn(s.ReturnKeywords.Span())
	}
	return NewBoundReturnStatements(expression)
}

func (b *Binder) BindIfStatement(s *syntax.IfStatement) Boundstatement {
	condition := b.BindExpressionAndCheckType(s.Condition, symbol.TypeBool)
	thenStatement := b.BindStatement(s.ThenStatement)
	var elseStatement Boundstatement
	if s.ElseClause != nil {
		elseStatement = b.BindStatement(s.ElseClause.ElseStatement)
	}

	return NewBoundIfStatements(condition, thenStatement, elseStatement)
}

func (b *Binder) BindBlockStatement(s *syntax.BlockStatement) Boundstatement {
	statements := make([]Boundstatement, 0)
	b.scope = NewBoundScope(b.scope)
	for _, statement := range s.Statements {
		statements = append(statements, b.BindStatement(statement))
	}
	b.scope = b.scope.Parent
	return NewBoundBlockStatement(statements)
}

func (b *Binder) BindVariableDeclaration(s *syntax.VariableDeclarationSyntax) Boundstatement {

	isReadOnly := s.Keyword.Kind() == syntax.SyntaxKindLetKeywords
	tp := b.BindTypeClause(s.TypeClause)
	initializer := b.BindExpression(s.Initializer, false)

	vType := tp
	if vType == nil {
		vType = initializer.Type()
	} else {
		if vType != initializer.Type() {
			b.Diagnostics.Report(s.TypeClause.Identifier.Span(), "Not support auto convert type")
			return NewBoundErrorExpression()
		}
	}
	variable := b.BindVariableSymbol(s.Identifier, isReadOnly, vType)

	return NewBoundVariableDeclaration(variable, initializer)
}

func (b *Binder) BindExpressionStatement(es *syntax.ExpressStatement) Boundstatement {
	expression := b.BindExpressionInternal(es.Express)
	return NewBoundExpressStatements(expression)
}

func (b *Binder) BindLiteralExpress(express *syntax.LiteralExpress) BoundExpression {
	value := express.Value
	if value == nil {
		value = int64(0)
	}
	return NewBoundLiteralExpression(value)
}

func (b *Binder) BindNameExpress(express *syntax.NameExpress) BoundExpression {
	if express.Identifier.Missing {
		return NewBoundErrorExpression()
	}
	name := express.Identifier.Text
	variable, has := b.scope.TryLookupVariable(name)
	if !has {
		b.Diagnostics.UndefinedName(express.Identifier.Span(), name)
		return NewBoundErrorExpression()
	}
	return NewBoundVariableExpression(variable)
}

func (b *Binder) BindAssignmentExpress(express *syntax.AssignmentExpress) BoundExpression {

	name := express.Identifier.Text
	boundExpress := b.BindExpression(express.Express, false)
	variable, has := b.scope.TryLookupVariable(name)
	if !has {
		b.Diagnostics.UndefinedName(express.Identifier.Span(), name)
		return boundExpress
	}
	if variable.IsReadOnly() {
		b.Diagnostics.CannotAssign(express.Identifier.Span(), name)
		return boundExpress
	}

	if boundExpress.Type() != variable.GetType() {
		b.Diagnostics.CannotConvert(express.Identifier.Span(), boundExpress.Type(), variable.GetType())
		return boundExpress
	}

	return NewBoundAssignmentExpression(variable, boundExpress)
}

func (b *Binder) BindUnaryExpress(express *syntax.UnaryExpress) BoundExpression {
	boundOperand := b.BindExpression(express.Operand, false)

	if boundOperand.Type() == symbol.TypeError {
		return NewBoundErrorExpression()
	}
	boundOperator := BindBoundUnaryOperator(express.Operator.Kind(), boundOperand.Type())

	if boundOperator == nil {
		b.Diagnostics.UndefinedUnaryOperator(express.Operator.Span(), express.Operator.Text, boundOperand.Type())
		return NewBoundErrorExpression()
	}

	return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder) BindBinaryOperator(express *syntax.BinaryExpress) BoundExpression {
	boundLeft := b.BindExpression(express.Left, false)
	boundRight := b.BindExpression(express.Right, false)

	if boundLeft.Type() == symbol.TypeError || boundRight.Type() == symbol.TypeError {
		return NewBoundErrorExpression()
	}

	boundOperator := BindBoundBinaryOperator(express.Operator.Kind(), boundLeft.Type(), boundRight.Type())

	if boundOperator == nil {
		b.Diagnostics.UndefinedBinaryOperator(express.Operator.Span(), express.Operator.Text, boundLeft.Type(), boundRight.Type())
		return NewBoundErrorExpression()
	}

	return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}

func (b *Binder) BindVariableSymbol(identifier syntax.SyntaxToken, isReadOnly bool, tp *symbol.TypeSymbol) symbol.VariableSymbol {
	name := identifier.Text
	if name == "" {
		name = "?"
	}

	var variable symbol.VariableSymbol
	if b.function == nil {
		variable = symbol.NewGlobalVariableSymbol(name, isReadOnly, tp)
	} else {
		variable = symbol.NewLocalVariableSymbol(name, isReadOnly, tp)
	}

	if !identifier.Missing && !b.scope.TryDeclareVariable(variable) {
		b.Diagnostics.VariableAlreadyDeclared(identifier.Span(), name)
	}

	return variable
}

func (b *Binder) BindCallExpression(express *syntax.CallExpress) BoundExpression {
	boundArguments := []BoundExpression{}
	arguments := express.Params.List()
	for _, arg := range arguments {
		boundArg := b.BindExpression(arg, false)
		boundArguments = append(boundArguments, boundArg)
	}

	function, has := b.scope.TryLookupFunction(express.Identifier.Text)
	if !has {
		b.Diagnostics.UndefinedFunction(express.Identifier.Span(), express.Identifier.Text)
		return NewBoundErrorExpression()
	}

	if len(function.Parameter) != len(boundArguments) {
		b.Diagnostics.WrongArgumentNumber(syntax.SyntaxNodeSpan(express),
			function.GetName(), len(function.Parameter), len(boundArguments))

		return NewBoundErrorExpression()
	}

	for i := 0; i < len(boundArguments); i++ {
		arg := boundArguments[i]
		param := function.Parameter[i]
		if arg.Type() != param.Type {
			b.Diagnostics.WrongArgumentType(syntax.SyntaxNodeSpan(arguments[i]),
				param.GetName(), param.Type, arg.Type())
			return NewBoundErrorExpression()
		}
	}

	return NewBoundcallExpression(function, boundArguments)
}

func (b *Binder) BindUnitExpression(express *syntax.UnitExpress) BoundExpression {
	return NewBoundUnitExpression(express)
}

func (b *Binder) BindFunctionDeclaration(s *syntax.FunctionDeclarationSyntax) {
	parameters := make([]*symbol.ParameterSymbol, 0)

	existParameterNames := map[string]struct{}{}

	for _, p := range s.ParameterList.List() {
		parameter := p.(*syntax.ParameterSyntax)
		name := parameter.Identifier.Text
		tp := b.BindTypeClause(parameter.Type)

		if _, has := existParameterNames[name]; has {
			b.Diagnostics.ParameterAlreadyDeclared(parameter.Identifier.Span(), name)
		} else {
			ps := symbol.NewParameterSymbol(name, tp)
			parameters = append(parameters, ps)
		}
	}

	resType := b.BindTypeClause(s.Type)
	if resType == nil {
		resType = symbol.TypeUnit
	}

	function := symbol.NewFunctionSymbol(s.Identifier.Text, parameters, resType, s)

	ok := b.scope.TryDeclareFunction(function)
	if !ok {
		b.Diagnostics.SymbolAlreadyDeclared(s.Identifier.Span(), function.GetName())
	}
}

func (b *Binder) BindTypeClause(s *syntax.TypeClauseSyntax) *symbol.TypeSymbol {
	if s == nil {
		return nil
	}

	tp := lookupType(s.Identifier.Text)
	if tp == nil {
		b.Diagnostics.UndefinedType(s.Identifier.Span(), s.Identifier.Text)
	}
	return tp
}

func lookupType(name string) *symbol.TypeSymbol {
	switch name {
	case "bool":
		return symbol.TypeBool
	case "int":
		return symbol.TypeInt
	case "string":
		return symbol.TypeString
	default:
		return symbol.TypeError
	}
}
