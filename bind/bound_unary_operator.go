package bind

import (
	"fmt"
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

func (b *BoundUnaryOperator) String() string {
	if str, has := operandTypeNameMap[b.Kind]; has {
		return str
	}
	return fmt.Sprintf("Unexcepted_BoundUnaryOperatorKind_%d", b)
}

var operandTypeNameMap = map[BoundUnaryOperatorKind]string{
	BoundUnaryOperatorKindIdentity:              "Identity",
	BoundUnaryOperatorKindNegation:              "Negation",
	BoundUnaryOperatorKindLogicalNegation:       "LogicalNegation",
	BoundUnaryOperatorKindBitwiseOnesComplement: "BitwiseOnesComplement",
}

func BindBoundUnaryOperator(syntaxKind syntax.SyntaxKind, operandType reflect.Kind) *BoundUnaryOperator {
	if syntaxKind == syntax.SyntaxKindPlusToken && operandType == reflect.Int64 {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindIdentity, reflect.Int64)
	} else if syntaxKind == syntax.SyntaxKindMinusToken && operandType == reflect.Int64 {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindNegation, reflect.Int64)
	} else if syntaxKind == syntax.SyntaxKindBangToken && operandType == reflect.Bool {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindLogicalNegation, reflect.Bool)
	} else if syntaxKind == syntax.SyntaxKindTildeToken && operandType == reflect.Int64 {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindBitwiseOnesComplement, reflect.Int64)
	} else {
		return nil
	}
}

type BoundUnaryOperatorKind int

const (
	BoundUnaryOperatorKindIdentity BoundUnaryOperatorKind = iota
	BoundUnaryOperatorKindNegation
	BoundUnaryOperatorKindLogicalNegation
	BoundUnaryOperatorKindBitwiseOnesComplement
)
