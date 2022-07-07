package bind

import "reflect"

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
