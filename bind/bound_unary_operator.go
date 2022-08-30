package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

type BoundUnaryOperator struct {
	syntaxKind  syntax.SyntaxKind
	Kind        BoundUnaryOperatorKind
	operandType *symbol.TypeSymbol
}

func newBoundUnaryOperator(syntaxKind syntax.SyntaxKind, kind BoundUnaryOperatorKind, operandType *symbol.TypeSymbol) *BoundUnaryOperator {
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
	return fmt.Sprintf("Unexcepted_BoundUnaryOperatorKind_%d", b.Kind)
}

var operandTypeNameMap = map[BoundUnaryOperatorKind]string{
	BoundUnaryOperatorKindIdentity:              "Identity",
	BoundUnaryOperatorKindNegation:              "Negation",
	BoundUnaryOperatorKindLogicalNegation:       "LogicalNegation",
	BoundUnaryOperatorKindBitwiseOnesComplement: "BitwiseOnesComplement",
}

func BindBoundUnaryOperator(syntaxKind syntax.SyntaxKind, operandType *symbol.TypeSymbol) *BoundUnaryOperator {
	if syntaxKind == syntax.SyntaxKindPlusToken && operandType == symbol.TypeInt {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindIdentity, symbol.TypeInt)
	} else if syntaxKind == syntax.SyntaxKindMinusToken && operandType == symbol.TypeInt {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindNegation, symbol.TypeInt)
	} else if syntaxKind == syntax.SyntaxKindBangToken && operandType == symbol.TypeBool {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindLogicalNegation, symbol.TypeBool)
	} else if syntaxKind == syntax.SyntaxKindTildeToken && operandType == symbol.TypeInt {
		return newBoundUnaryOperator(syntaxKind, BoundUnaryOperatorKindBitwiseOnesComplement, symbol.TypeInt)
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
