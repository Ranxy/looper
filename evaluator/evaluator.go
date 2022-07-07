package evaluator

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
)

func Evaluate(node bind.BoundExpression) any {
	if n, ok := node.(*bind.BoundLiteralExpression); ok {
		return n.Value
	}
	if n, ok := node.(*bind.BoundUnaryExpression); ok {
		operand := Evaluate(n.Operand)
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
		left := Evaluate(n.Left)
		right := Evaluate(n.Right)

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
