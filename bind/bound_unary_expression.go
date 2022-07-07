package bind

import "reflect"

type BoundUnaryExpression struct {
	Op      *BoundUnaryOperator
	Operand BoundExpression
}

func NewBoundUnaryExpression(op *BoundUnaryOperator, operand BoundExpression) *BoundUnaryExpression {
	return &BoundUnaryExpression{
		Op:      op,
		Operand: operand,
	}
}
func (b *BoundUnaryExpression) Type() reflect.Kind {
	return b.Operand.Type()
}
func (b *BoundUnaryExpression) Kind() BoundNodeKind {
	return BoundNodeKindUnaryExpress
}
