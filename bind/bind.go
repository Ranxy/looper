package bind

import (
	"container/list"
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/syntax"
)

type Binder struct {
	Diagnostics *diagnostic.DiagnosticBag
	scope       *BoundScope
}

func NewBinder(parent *BoundScope) *Binder {
	return &Binder{
		Diagnostics: diagnostic.NewDiagnostics(),
		scope:       NewBoundScope(parent),
	}
}

func BindGlobalScope(previous *BoundGlobalScope, s *syntax.CompliationUnit) *BoundGlobalScope {
	parentScope := CreateParentScope(previous)
	binder := NewBinder(parentScope)
	expression := binder.BindStatement(s.Statement)
	variables := binder.scope.GetDeclareVariables()
	diagnostics := binder.Diagnostics

	if previous != nil {
		diagnostics = diagnostics.Merge(previous.Diagnostic)
	}

	return NewBoundGlobalScope(previous, diagnostics, variables, expression)
}

func CreateParentScope(previous *BoundGlobalScope) *BoundScope {

	stack := list.New()

	for previous != nil {
		stack.PushBack(previous)
		previous = previous.Previous
	}

	var parent *BoundScope

	for stack.Len() > 0 {
		pvi := stack.Back()
		if pvi != nil {
			stack.Remove(pvi)
		}
		pv := pvi.Value.(*BoundGlobalScope)
		scope := NewBoundScope(parent)
		for _, v := range pv.Variables {
			scope.TryDeclare(v)
		}
		parent = scope
	}
	return parent
}

func (b *Binder) BindExpressionAndCheckType(express syntax.Express, targetType reflect.Kind) BoundExpression {
	result := b.BindExpression(express)
	if result.Type() != targetType {
		b.Diagnostics.CannotConvert(syntax.SyntaxNodeSpan(express), result.Type(), targetType)
	}
	return result
}

func (b *Binder) BindExpression(express syntax.Express) BoundExpression {
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
		return b.BindExpression(express.(*syntax.ParenthesisExpress).Expr)
	case syntax.SyntaxKindAssignmentExpress:
		return b.BindAssignmentExpress(express.(*syntax.AssignmentExpress))
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
	case syntax.SyntaxKindExpressStatement:
		return b.BindExpressionStatement(s.(*syntax.ExpressStatement))
	default:
		panic(fmt.Sprintf("Unexceped syntax %s", s.Kind()))
	}
}

func (b *Binder) BindForStatement(s *syntax.ForStatement) Boundstatement {
	initCond := b.BindStatement(s.InitCondition)
	endCond := b.BindExpressionAndCheckType(s.EndCondition, reflect.Bool)
	updateCond := b.BindStatement(s.UpdateCondition)
	body := b.BindStatement(s.Body)
	return NewBoundForStatements(initCond, endCond, updateCond, body)
}

func (b *Binder) BindWhileStatement(s *syntax.WhileStatement) Boundstatement {
	condition := b.BindExpressionAndCheckType(s.Condition, reflect.Bool)
	body := b.BindStatement(s.Body)
	return NewBoundWhileStatements(condition, body)
}

func (b *Binder) BindIfStatement(s *syntax.IfStatement) Boundstatement {
	condition := b.BindExpressionAndCheckType(s.Condition, reflect.Bool)
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
	name := s.Identifier.Text
	isReadOnly := s.Keyword.Kind() == syntax.SyntaxKindLetKeywords
	initializer := b.BindExpression(s.Initializer)
	variable := syntax.NewVariableSymbol(name, isReadOnly, initializer.Type())

	if !b.scope.TryDeclare(variable) {
		b.Diagnostics.VariableAlreadyDeclared(s.Identifier.Span(), name)
	}

	return NewBoundVariableDeclaration(variable, initializer)
}

func (b *Binder) BindExpressionStatement(es *syntax.ExpressStatement) Boundstatement {
	expression := b.BindExpression(es.Express)
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
	name := express.Identifier.Text
	if name == "" {
		return NewBoundLiteralExpression(0)
	}
	variable, has := b.scope.TryLookup(name)
	if !has {
		b.Diagnostics.UndefinedName(express.Identifier.Span(), name)
		return NewBoundLiteralExpression(int64(0))
	}
	return NewBoundVariableExpression(variable)
}

func (b *Binder) BindAssignmentExpress(express *syntax.AssignmentExpress) BoundExpression {

	name := express.Identifier.Text
	boundExpress := b.BindExpression(express.Express)
	variable, has := b.scope.TryLookup(name)
	if !has {
		b.Diagnostics.UndefinedName(express.Identifier.Span(), name)
		return boundExpress
	}
	if variable.IsReadOnly {
		b.Diagnostics.CannotAssign(express.Identifier.Span(), name)
		return boundExpress
	}

	if boundExpress.Type() != variable.Type {
		b.Diagnostics.CannotConvert(express.Identifier.Span(), boundExpress.Type(), variable.Type)
		return boundExpress
	}

	return NewBoundAssignmentExpression(variable, boundExpress)
}

func (b *Binder) BindUnaryExpress(express *syntax.UnaryExpress) BoundExpression {
	boundOperand := b.BindExpression(express.Operand)
	boundOperator := BindBoundUnaryOperator(express.Operator.Kind(), boundOperand.Type())

	if boundOperator == nil {
		b.Diagnostics.UndefinedUnaryOperator(express.Operator.Span(), express.Operator.Text, boundOperand.Type())
		return boundOperand
	}

	return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder) BindBinaryOperator(express *syntax.BinaryExpress) BoundExpression {
	boundLeft := b.BindExpression(express.Left)
	boundRight := b.BindExpression(express.Right)
	boundOperator := BindBoundBinaryOperator(express.Operator.Kind(), boundLeft.Type(), boundRight.Type())

	if boundOperator == nil {
		b.Diagnostics.UndefinedBinaryOperator(express.Operator.Span(), express.Operator.Text, boundLeft.Type(), boundRight.Type())
		return boundLeft
	}

	return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}
