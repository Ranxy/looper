package bind

import (
	"fmt"
	"reflect"
)

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
func (b *BoundBinaryExpression) GetChildren() []BoundNode {
	return []BoundNode{b.Left, b.Right}
}

func (b *BoundBinaryExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Op}
}
