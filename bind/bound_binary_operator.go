package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

type BoundBinaryOperator struct {
	SyntaxKind syntax.SyntaxKind
	Kind       BoundBinaryOperatorKind
	LeftType   *symbol.TypeSymbol
	RightType  *symbol.TypeSymbol
	Type       *symbol.TypeSymbol
}

func BindBoundBinaryOperator(syntaxKind syntax.SyntaxKind, leftType, rightType *symbol.TypeSymbol) *BoundBinaryOperator {
	f, has := binaryMatchParam[binaryMatchType{
		SyntaxKind: syntaxKind,
		LeftType:   leftType,
		RightType:  rightType,
	}]
	if !has {
		return nil
	}
	return f()
}

func (b *BoundBinaryOperator) String() string {
	return b.Kind.String()
}

type binaryMatchType struct {
	SyntaxKind syntax.SyntaxKind
	LeftType   *symbol.TypeSymbol
	RightType  *symbol.TypeSymbol
}

var binaryMatchParam = map[binaryMatchType]func() *BoundBinaryOperator{
	{syntax.SyntaxKindPlusToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPlusToken, BoundBinaryKindAddition, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindMinusToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindMinusToken, BoundBinaryKindSubtraction, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindStarToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindStarToken, BoundBinaryKindMultiplication, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindSlashToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindSlashToken, BoundBinaryKindDivision, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindEqualEqualToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindEqualEqualToken, BoundBinaryKindEquals, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},
	{syntax.SyntaxKindBangEqualToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindBangEqualToken, BoundBinaryKindNotEquals, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},
	{syntax.SyntaxKindAmpersandAmpersandToken, symbol.TypeBool, symbol.TypeBool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindAmpersandAmpersandToken, BoundBinaryKindLogicalAnd, symbol.TypeBool, symbol.TypeBool, symbol.TypeBool}
	},
	{syntax.SyntaxKindPipePileToken, symbol.TypeBool, symbol.TypeBool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPipePileToken, BoundBinaryKindLogicalOr, symbol.TypeBool, symbol.TypeBool, symbol.TypeBool}
	},
	{syntax.SyntaxKindEqualEqualToken, symbol.TypeBool, symbol.TypeBool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindEqualEqualToken, BoundBinaryKindEquals, symbol.TypeBool, symbol.TypeBool, symbol.TypeBool}
	},
	{syntax.SyntaxKindBangEqualToken, symbol.TypeBool, symbol.TypeBool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindBangEqualToken, BoundBinaryKindNotEquals, symbol.TypeBool, symbol.TypeBool, symbol.TypeBool}
	},
	{syntax.SyntaxKindLessToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindLessToken, BoundBinaryKindLess, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},
	{syntax.SyntaxKindLessEqualToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindLessEqualToken, BoundBinaryKindLessEqual, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},
	{syntax.SyntaxKindGreatToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindGreatToken, BoundBinaryKindGreat, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},
	{syntax.SyntaxKindGreatEqualToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindGreatEqualToken, BoundBinaryKindGreatEqual, symbol.TypeInt, symbol.TypeInt, symbol.TypeBool}
	},

	//bitwise
	{syntax.SyntaxKindAmpersandToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindAmpersandToken, BoundBinaryKindBitwiseAnd, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindPipeToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPipeToken, BoundBinaryKindBitwiseOr, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},
	{syntax.SyntaxKindHatToken, symbol.TypeInt, symbol.TypeInt}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindHatToken, BoundBinaryKindBitwiseXor, symbol.TypeInt, symbol.TypeInt, symbol.TypeInt}
	},

	// string operator
	{syntax.SyntaxKindPlusToken, symbol.TypeString, symbol.TypeString}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPlusToken, BoundBinaryKindStringAdd, symbol.TypeString, symbol.TypeString, symbol.TypeString}
	},
	{syntax.SyntaxKindEqualEqualToken, symbol.TypeString, symbol.TypeString}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindEqualEqualToken, BoundBinaryKindStringEqual, symbol.TypeString, symbol.TypeString, symbol.TypeBool}
	},
	{syntax.SyntaxKindBangEqualToken, symbol.TypeString, symbol.TypeString}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindBangEqualToken, BoundBinaryKindStringNotEqual, symbol.TypeString, symbol.TypeString, symbol.TypeBool}
	},
}

type BoundBinaryOperatorKind int

const (
	BoundBinaryKindAddition BoundBinaryOperatorKind = iota
	BoundBinaryKindSubtraction
	BoundBinaryKindMultiplication
	BoundBinaryKindDivision
	BoundBinaryKindLogicalAnd
	BoundBinaryKindLogicalOr
	BoundBinaryKindBitwiseAnd
	BoundBinaryKindBitwiseOr
	BoundBinaryKindBitwiseXor
	BoundBinaryKindEquals
	BoundBinaryKindNotEquals
	BoundBinaryKindLess
	BoundBinaryKindLessEqual
	BoundBinaryKindGreat
	BoundBinaryKindGreatEqual
	BoundBinaryKindStringAdd
	BoundBinaryKindStringEqual
	BoundBinaryKindStringNotEqual
)

func (b BoundBinaryOperatorKind) String() string {
	str, has := boundBinaryKindStrMap[b]
	if has {
		return str
	} else {
		return fmt.Sprintf("Unexcepted_BoundBinaryOperatorKind_%d", b)
	}
}

var boundBinaryKindStrMap = map[BoundBinaryOperatorKind]string{
	BoundBinaryKindAddition:       "BoundBinaryKindAddition",
	BoundBinaryKindSubtraction:    "BoundBinaryKindSubtraction",
	BoundBinaryKindMultiplication: "BoundBinaryKindMultiplication",
	BoundBinaryKindDivision:       "BoundBinaryKindDivision",
	BoundBinaryKindLogicalAnd:     "BoundBinaryKindLogicalAnd",
	BoundBinaryKindLogicalOr:      "BoundBinaryKindLogicalOr",
	BoundBinaryKindBitwiseAnd:     "BoundBinaryKindBitwiseAnd",
	BoundBinaryKindBitwiseOr:      "BoundBinaryKindBitwiseOr",
	BoundBinaryKindBitwiseXor:     "BoundBinaryKindBitwiseXor",
	BoundBinaryKindEquals:         "BoundBinaryKindEquals",
	BoundBinaryKindNotEquals:      "BoundBinaryKindNotEquals",
	BoundBinaryKindLess:           "BoundBinaryKindLess",
	BoundBinaryKindLessEqual:      "BoundBinaryKindLessEqual",
	BoundBinaryKindGreat:          "BoundBinaryKindGreat",
	BoundBinaryKindGreatEqual:     "BoundBinaryKindGreatEqual",
	BoundBinaryKindStringAdd:      "BoundBinaryKindStringAdd",
}
