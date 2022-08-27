package bind

import (
	"fmt"
	"reflect"
)

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

func (b *BoundUnaryExpression) GetChildren() []BoundNode {
	return []BoundNode{b.Operand}

}
func (b *BoundUnaryExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Op}
}
