package evaluator

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
)

type Evaluater struct {
	root bind.BoundExpression
	vm   bind.VariableManage
}

func NewEvaluater(root bind.BoundExpression, vm bind.VariableManage) *Evaluater {
	return &Evaluater{
		root: root,
		vm:   vm,
	}
}

func (e *Evaluater) Evaluate() any {
	return e.EvaluateExpression(e.root)
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
	v := e.vm.GetValue(node.Variable.Name)
	if v == nil {
		panic(fmt.Sprintf("Undefined Variable %s", node.Variable.Name))
	}
	return v.Value
}

func (e *Evaluater) evaluateAssignmentExpression(node *bind.BoundAssignmentExpression) any {
	value := e.EvaluateExpression(node.Express)
	e.vm.Add(node.Variable, value)
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
	default:
		panic(fmt.Sprintf("Unexceped binary operator:%q and", node.Op))
	}
}
