package bind

import (
	"fmt"
	"reflect"
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
func (b *BoundLiteralExpression) Type() reflect.Kind {
	return reflect.TypeOf(b.Value).Kind()
}

func (b *BoundLiteralExpression) GetChildren() []BoundNode {
	return []BoundNode{}

}

func (b *BoundLiteralExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{&literalValue{v: b.Value}}
}
