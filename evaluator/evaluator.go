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
	if n, ok := node.(*bind.BoundLiteralExpression); ok {
		return n.Value
	}
	if n, ok := node.(*bind.BoundVariableExpression); ok {
		v := e.vm.GetValue(n.Variable.Name)
		if v == nil {
			panic(fmt.Sprintf("Undefined Variable %s", n.Variable.Name))
		}
		return v.Value
	}

	if n, ok := node.(*bind.BoundAssignmentExpression); ok {
		value := e.EvaluateExpression(n.Express)
		e.vm.Add(n.Variable, value)
		return value
	}
	if n, ok := node.(*bind.BoundUnaryExpression); ok {
		operand := e.EvaluateExpression(n.Operand)
		switch n.Op.Kind {
		case bind.BoundUnaryOperatorKindIdentity:
			return operand.(int64)
		case bind.BoundUnaryOperatorKindNegation:
			return -(operand.(int64))
		case bind.BoundUnaryOperatorKindLogicalNegation:
			return !(operand.(bool))
		default:
			panic(fmt.Sprintf("Unexceped unary operator:%v", n.Op))
		}
	}

	if n, ok := node.(*bind.BoundBinaryExpression); ok {
		left := e.EvaluateExpression(n.Left)
		right := e.EvaluateExpression(n.Right)

		switch n.Op.Kind {
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
			panic(fmt.Sprintf("Unexceped binary operator:%q and", n.Op))
		}
	}

	panic(fmt.Sprintf("Unexceped node %v", node))
}
