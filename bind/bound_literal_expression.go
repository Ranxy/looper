package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundLiteralExpression struct {
	Value any
}

func NewBoundLiteralExpression(value any) *BoundLiteralExpression {
	return &BoundLiteralExpression{
		Value: value,
	}
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind {
	return BoundNodeKindLiteralExpress
}
func (b *BoundLiteralExpression) Type() *symbol.TypeSymbol {
	switch b.Value.(type) {
	case int64:
		return symbol.TypeInt
	case bool:
		return symbol.TypeBool
	case string:
		return symbol.TypeString
	default:
		return symbol.TypeUnkonw
	}
}

func (b *BoundLiteralExpression) GetChildren() []BoundNode {
	return []BoundNode{}

}

func (b *BoundLiteralExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{&literalValue{v: b.Value}}
}
