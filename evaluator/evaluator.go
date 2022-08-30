package evaluator

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/symbol"
)

type Evaluater struct {
	root *bind.BoundBlockStatements
	vm   map[symbol.VariableSymbol]any

	lastValue any
}

func NewEvaluater(root *bind.BoundBlockStatements, vm map[symbol.VariableSymbol]any) *Evaluater {
	return &Evaluater{
		root: root,
		vm:   vm,
	}
}

func (e *Evaluater) Evaluate() any {
	labelToIndex := make(map[*bind.BoundLabel]int)
	for i := range e.root.Statement {
		if bls, ok := e.root.Statement[i].(*bind.LabelStatement); ok {
			labelToIndex[bls.Label] = i + 1
		}
	}
	index := 0
	for index < len(e.root.Statement) {
		s := e.root.Statement[index]
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
	e.vm[*node.Variable] = value
	e.lastValue = value
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
	default:
		panic(fmt.Sprintf("Unexceped node %v", node))
	}
}

func (e *Evaluater) evaluateLiteralExpression(node *bind.BoundLiteralExpression) any {
	return node.Value
}

func (e *Evaluater) evaluateVariableExpression(node *bind.BoundVariableExpression) any {
	v := e.vm[*node.Variable]
	if v == nil {
		panic(fmt.Sprintf("Undefined Variable %s", node.Variable.Name))
	}
	return v
}

func (e *Evaluater) evaluateAssignmentExpression(node *bind.BoundAssignmentExpression) any {
	value := e.EvaluateExpression(node.Express)
	e.vm[*node.Variable] = value
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

	default:
		panic(fmt.Sprintf("Unexceped binary operator:%q and", node.Op))
	}
}
