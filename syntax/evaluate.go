package syntax

import (
	"errors"
	"fmt"
	"reflect"
)

func Evaluate(node Express) (any, error) {
	if n, ok := node.(*LiteralExpress); ok {
		switch n.Literal.kind {
		case SyntaxKindNumberToken:
			return n.Literal.Value.(int64), nil
		case SyntaxKindTrueKeywords, SyntaxKindFalseKeywords:
			return n.Value.(bool), nil
		default:
			return nil, fmt.Errorf("Unsupport LiteralExpress %s", n.Kind())
		}
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
			return -v.(int64), nil
		case SyntaxKindBangToken:
			v, err := Evaluate(n.Operand)
			if err != nil {
				return 0, err
			}
			return !v.(bool), nil
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
			return left.(int64) + right.(int64), nil
		case SyntaxKindMinusToken:
			return left.(int64) - right.(int64), nil
		case SyntaxKindStarToken:
			return left.(int64) * right.(int64), nil
		case SyntaxKindSlashToken:
			return left.(int64) / right.(int64), nil

		case SyntaxKindAmpersandAmpersandToken:
			return left.(bool) && right.(bool), nil
		case SyntaxKindPipePileToken:
			return left.(bool) || right.(bool), nil

		case SyntaxKindEqualEqualToken:
			switch left.(type) {
			case int64:
				return left.(int64) == right.(int64), nil
			case bool:
				return left.(bool) == right.(bool), nil
			default:
				return 0, errors.New(fmt.Sprintf("SyntaxKindEqualEqualToken Left Type %v And Right Type %s Not equal ", reflect.TypeOf(left), reflect.TypeOf(right)))
			}
		case SyntaxKindBangEqualToken:
			switch left.(type) {
			case int64:
				return left.(int64) != right.(int64), nil
			case bool:
				return left.(bool) != right.(bool), nil
			default:
				return 0, errors.New(fmt.Sprintf("SyntaxKindBangEqualToken Left Type %v And Right Type %s Not equal ", reflect.TypeOf(left), reflect.TypeOf(right)))
			}
		default:
			return 0, errors.New(fmt.Sprintf("BinaryOperator Kind %s not found ", n.Operator.Kind()))
		}

	default:
		return 0, errors.New(fmt.Sprintf("Unexceped Node %s ", n.Kind()))
	}
}
