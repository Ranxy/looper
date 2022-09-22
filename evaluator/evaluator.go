package evaluator

import (
	"container/list"
	"fmt"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/bind/program"
	"github.com/Ranxy/looper/buildin"
	"github.com/Ranxy/looper/symbol"
)

type vStore map[symbol.VariableSymbol]any

type Evaluater struct {
	program *program.BoundProgram
	global  vStore
	local   *list.List //stack[map[symbol.VariableSymbol]any]

	lastValue any
}

func NewEvaluater(program *program.BoundProgram, vm map[symbol.VariableSymbol]any) *Evaluater {
	return &Evaluater{
		program:   program,
		global:    vm,
		local:     &list.List{},
		lastValue: nil,
	}
}
func (e *Evaluater) Evaluate() any {
	return e.EvaluateStatement(e.program.Statement)
}

func (e *Evaluater) EvaluateStatement(body *bind.BoundBlockStatements) any {
	if body == nil {
		return nil
	}
	labelToIndex := make(map[*bind.BoundLabel]int)
	for i := range body.Statement {
		if bls, ok := body.Statement[i].(*bind.LabelStatement); ok {
			labelToIndex[bls.Label] = i + 1
		}
	}
	index := 0
	for index < len(body.Statement) {
		s := body.Statement[index]
		switch s.Kind() {
		case bind.BoundNodeKindVariableDeclaration:
			e.EvaluateVariableDeclaration(s.(*bind.BoundVariableDeclaration))
			index += 1
		case bind.BoundNodeKindExpressionStatement:
			e.EvaluateExpressionStatement(s.(*bind.BoundExpressStatements))
			index += 1
		case bind.BoundNodeKindGotoStatement:
			gts := s.(*bind.GotoStatement)
			index = labelToIndex[gts.Label]
		case bind.BoundNodeKindConditionalGotoStatement:
			cgts := s.(*bind.ConditionalGotoStatement)
			conditon := e.EvaluateExpression(cgts.Condition).(bool)
			if conditon && !cgts.JumpIfFalse || !conditon && cgts.JumpIfFalse {
				index = labelToIndex[cgts.Label]
			} else {
				index += 1
			}
		case bind.BoundNodeKindLabelStatement:
			index += 1
		default:
			panic(fmt.Sprintf("Uncxcepted node %v", s.Kind()))
		}
	}

	return e.lastValue
}

func (e *Evaluater) EvaluateVariableDeclaration(node *bind.BoundVariableDeclaration) {
	value := e.EvaluateExpression(node.Initializer)
	e.lastValue = value
	e.Assign(node.Variable, value)
}

func (e *Evaluater) EvaluateExpressionStatement(node *bind.BoundExpressStatements) {
	e.lastValue = e.EvaluateExpression(node.Express)
}

func (e *Evaluater) EvaluateExpression(node bind.BoundExpression) any {

	switch n := node.(type) {
	case *bind.BoundLiteralExpression:
		return e.evaluateLiteralExpression(n)
	case *bind.BoundVariableExpression:
		return e.evaluateVariableExpression(n)
	case *bind.BoundAssignmentExpression:
		return e.evaluateAssignmentExpression(n)
	case *bind.BoundUnaryExpression:
		return e.evaluateUnaryExpression(n)
	case *bind.BoundBinaryExpression:
		return e.evaluateBinaryExpression(n)
	case *bind.BoundCallExpression:
		return e.evaluateCallExpression(n)
	default:
		panic(fmt.Sprintf("Unexceped node %v", node))
	}
}

func (e *Evaluater) evaluateLiteralExpression(node *bind.BoundLiteralExpression) any {
	return node.Value
}

func (e *Evaluater) evaluateVariableExpression(node *bind.BoundVariableExpression) any {

	if node.Variable.Kind() == symbol.SymbolKindGlobalVariable {
		return e.global[node.Variable]
	} else {
		local := e.local.Back().Value.(vStore)
		return local[node.Variable]
	}
}

func (e *Evaluater) evaluateAssignmentExpression(node *bind.BoundAssignmentExpression) any {
	value := e.EvaluateExpression(node.Express)
	e.Assign(node.Variable, value)
	return value
}

func (e *Evaluater) evaluateUnaryExpression(node *bind.BoundUnaryExpression) any {
	operand := e.EvaluateExpression(node.Operand)
	switch node.Op.Kind {
	case bind.BoundUnaryOperatorKindIdentity:
		return operand.(int64)
	case bind.BoundUnaryOperatorKindNegation:
		return -(operand.(int64))
	case bind.BoundUnaryOperatorKindLogicalNegation:
		return !(operand.(bool))
	case bind.BoundUnaryOperatorKindBitwiseOnesComplement:
		return ^(operand.(int64))
	default:
		panic(fmt.Sprintf("Unexceped unary operator:%v", node.Op))
	}
}

func (e *Evaluater) evaluateBinaryExpression(node *bind.BoundBinaryExpression) any {
	left := e.EvaluateExpression(node.Left)
	right := e.EvaluateExpression(node.Right)

	switch node.Op.Kind {
	case bind.BoundBinaryKindAddition:
		return left.(int64) + right.(int64)
	case bind.BoundBinaryKindSubtraction:
		return left.(int64) - right.(int64)
	case bind.BoundBinaryKindMultiplication:
		return left.(int64) * right.(int64)
	case bind.BoundBinaryKindDivision:
		return left.(int64) / right.(int64)
	case bind.BoundBinaryKindLogicalAnd:
		return left.(bool) && right.(bool)
	case bind.BoundBinaryKindLogicalOr:
		return left.(bool) || right.(bool)
	case bind.BoundBinaryKindEquals:
		return left == right
	case bind.BoundBinaryKindNotEquals:
		return left != right
	case bind.BoundBinaryKindLess:
		return left.(int64) < right.(int64)
	case bind.BoundBinaryKindLessEqual:
		return left.(int64) <= right.(int64)
	case bind.BoundBinaryKindGreat:
		return left.(int64) > right.(int64)
	case bind.BoundBinaryKindGreatEqual:
		return left.(int64) >= right.(int64)
	case bind.BoundBinaryKindBitwiseAnd:
		return left.(int64) & right.(int64)
	case bind.BoundBinaryKindBitwiseOr:
		return left.(int64) | right.(int64)
	case bind.BoundBinaryKindBitwiseXor:
		return left.(int64) ^ right.(int64)
	case bind.BoundBinaryKindStringAdd:
		return left.(string) + right.(string)
	case bind.BoundBinaryKindStringEqual:
		return left.(string) == right.(string)
	case bind.BoundBinaryKindStringNotEqual:
		return left.(string) != right.(string)
	default:
		panic(fmt.Sprintf("Unexceped binary operator:%q and", node.Op))
	}
}

func (e *Evaluater) evaluateCallExpression(node *bind.BoundCallExpression) any {
	if node.Function == buildin.FunctionPrint {
		msg := e.EvaluateExpression(node.Arguments[0]).(string)
		buildin.FunctionPrintImpl(msg)
		return nil
	} else if node.Function == buildin.FunctionInputStr {
		return buildin.FunctionInputStrImpl()
	} else if node.Function == buildin.FunctionRnd {
		max := e.EvaluateExpression(node.Arguments[0]).(int64)
		return buildin.FunctionRndImpl(max)
	} else {
		local := vStore{}

		for i, arg := range node.Arguments {
			parameter := node.Function.Parameter[i]
			value := e.EvaluateExpression(arg)
			local[parameter] = value
		}
		e.local.PushBack(local)

		statement := e.program.Functions[node.Function]
		result := e.EvaluateStatement(statement)
		e.local.Remove(e.local.Back())
		return result
	}
}

func (e *Evaluater) Assign(variable symbol.VariableSymbol, value any) {
	if variable.Kind() == symbol.SymbolKindGlobalVariable {
		e.global[variable] = value
	} else {
		lmr := e.local.Back()
		lm := lmr.Value.(vStore)
		lm[variable] = value
	}
}
