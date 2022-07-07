package bind

import (
	"reflect"

	"github.com/Ranxy/looper/syntax"
)

type BoundUnaryOperator struct {
	syntaxKind  syntax.SyntaxKind
	Kind        BoundUnaryOperatorKind
	operandType reflect.Kind
}

func newBoundUnaryOperator(syntaxKind syntax.SyntaxKind, kind BoundUnaryOperatorKind, operandType reflect.Kind) *BoundUnaryOperator {
	return &BoundUnaryOperator{
		syntaxKind:  syntaxKind,
		Kind:        kind,
		operandType: operandType,
	}
}

func BindBoundUnaryOperator(syntaxKind syntax.SyntaxKind, operandType reflect.Kind) *BoundUnaryOperator {
	if syntaxKind == syntax.SyntaxKindPlusToken && operandType == reflect.Int64 {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindIdentity, reflect.Int64)
	} else if syntaxKind == syntax.SyntaxKindMinusToken && operandType == reflect.Int64 {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindNegation, reflect.Int64)
	} else if syntaxKind == syntax.SyntaxKindBangToken && operandType == reflect.Bool {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindLogicalNegation, reflect.Bool)
	} else {
		return nil
	}
}

type BoundUnaryOperatorKind int

const (
	BoundUnaryOperatorKindIdentity BoundUnaryOperatorKind = iota
	BoundUnaryOperatorKindNegation
	BoundUnaryOperatorKindLogicalNegation
)
