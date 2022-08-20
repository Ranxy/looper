package bind

import (
	"reflect"

	"github.com/Ranxy/looper/syntax"
)

type BoundBinaryOperator struct {
	SyntaxKind syntax.SyntaxKind
	Kind       BoundBinaryOperatorKind
	LeftType   reflect.Kind
	RightType  reflect.Kind
	Type       reflect.Kind
}

func BindBoundBinaryOperator(syntaxKind syntax.SyntaxKind, leftType, rightType reflect.Kind) *BoundBinaryOperator {
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

type binaryMatchType struct {
	SyntaxKind syntax.SyntaxKind
	LeftType   reflect.Kind
	RightType  reflect.Kind
}

var binaryMatchParam = map[binaryMatchType]func() *BoundBinaryOperator{
	{syntax.SyntaxKindPlusToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPlusToken, BoundBinaryKindAddition, reflect.Int64, reflect.Int64, reflect.Int64}
	},
	{syntax.SyntaxKindMinusToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindMinusToken, BoundBinaryKindSubtraction, reflect.Int64, reflect.Int64, reflect.Int64}
	},
	{syntax.SyntaxKindStarToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindStarToken, BoundBinaryKindMultiplication, reflect.Int64, reflect.Int64, reflect.Int64}
	},
	{syntax.SyntaxKindSlashToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindSlashToken, BoundBinaryKindDivision, reflect.Int64, reflect.Int64, reflect.Int64}
	},
	{syntax.SyntaxKindEqualEqualToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindEqualEqualToken, BoundBinaryKindEquals, reflect.Int64, reflect.Int64, reflect.Bool}
	},
	{syntax.SyntaxKindBangEqualToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindBangEqualToken, BoundBinaryKindNotEquals, reflect.Int64, reflect.Int64, reflect.Bool}
	},
	{syntax.SyntaxKindAmpersandAmpersandToken, reflect.Bool, reflect.Bool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindAmpersandAmpersandToken, BoundBinaryKindLogicalAnd, reflect.Bool, reflect.Bool, reflect.Bool}
	},
	{syntax.SyntaxKindPipePileToken, reflect.Bool, reflect.Bool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindPipePileToken, BoundBinaryKindLogicalOr, reflect.Bool, reflect.Bool, reflect.Bool}
	},
	{syntax.SyntaxKindEqualEqualToken, reflect.Bool, reflect.Bool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindEqualEqualToken, BoundBinaryKindEquals, reflect.Bool, reflect.Bool, reflect.Bool}
	},
	{syntax.SyntaxKindBangEqualToken, reflect.Bool, reflect.Bool}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindBangEqualToken, BoundBinaryKindNotEquals, reflect.Bool, reflect.Bool, reflect.Bool}
	},
	{syntax.SyntaxKindLessToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindLessToken, BoundBinaryKindLess, reflect.Int64, reflect.Int64, reflect.Bool}
	},
	{syntax.SyntaxKindLessEqualToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindLessEqualToken, BoundBinaryKindLessEqual, reflect.Int64, reflect.Int64, reflect.Bool}
	},
	{syntax.SyntaxKindGreatToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindGreatToken, BoundBinaryKindGreat, reflect.Int64, reflect.Int64, reflect.Bool}
	},
	{syntax.SyntaxKindGreatEqualToken, reflect.Int64, reflect.Int64}: func() *BoundBinaryOperator {
		return &BoundBinaryOperator{syntax.SyntaxKindGreatEqualToken, BoundBinaryKindGreatEqual, reflect.Int64, reflect.Int64, reflect.Bool}
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
	BoundBinaryKindEquals
	BoundBinaryKindNotEquals
	BoundBinaryKindLess
	BoundBinaryKindLessEqual
	BoundBinaryKindGreat
	BoundBinaryKindGreatEqual
)
