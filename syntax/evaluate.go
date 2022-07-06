package syntax

import (
	"errors"
	"fmt"
)

func Evaluate(node Express) (int64, error) {
	if n, ok := node.(*LiteralExpress); ok {
		return n.Literal.Value.(int64), nil
	}

	switch n := node.(type) {
	case *LiteralExpress:
		return n.Literal.Value.(int64), nil
	case *ParenthesisExpress:
		return Evaluate(n.Expr)
	case *UnaryExpress:
		switch n.Operator.Kind() {
		case SyntaxKindPlusToken:
			return Evaluate(n.Operand)
		case SyntaxKindMinusToken:
			v, err := Evaluate(n.Operand)
			if err != nil {
				return 0, err
			}
			return -v, nil
		default:
			return 0, errors.New(fmt.Sprintf("UnaryOperator Kind %s not found ", n.Operator.Kind()))
		}
	case *BinaryExpress:
		left, err := Evaluate(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := Evaluate(n.Right)
		if err != nil {
			return 0, err
		}
		switch n.Operator.Kind() {
		case SyntaxKindPlusToken:
			return left + right, nil
		case SyntaxKindMinusToken:
			return left - right, nil
		case SyntaxKindStarToken:
			return left * right, nil
		case SyntaxKindSlashToken:
			return left / right, nil
		default:
			return 0, errors.New(fmt.Sprintf("BinaryOperator Kind %s not found ", n.Operator.Kind()))
		}

	default:
		return 0, errors.New(fmt.Sprintf("Unexceped Node %s ", n.Kind()))
	}
}
