package bind

import "reflect"

type BoundBinaryExpression struct {
	Left  BoundExpression
	Op    *BoundBinaryOperator
	Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, op *BoundBinaryOperator, right BoundExpression) *BoundBinaryExpression {
	return &BoundBinaryExpression{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (b *BoundBinaryExpression) Type() reflect.Kind {
	return b.Op.Type
}
func (b *BoundBinaryExpression) Kind() BoundNodeKind {
	return BoundNodeKindBinaryExpress
}
