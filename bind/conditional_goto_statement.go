package bind

import (
	"reflect"

	"github.com/Ranxy/looper/label"
)

type ConditionalGotoStatement struct {
	Label       *label.LabelSymbol
	Condition   BoundExpression
	JumpIfFalse bool
}

func NewConditionalGotoSymbol(label *label.LabelSymbol) *ConditionalGotoStatement {
	return &ConditionalGotoStatement{
		Label: label,
	}
}

func (b *ConditionalGotoStatement) Kind() BoundNodeKind {
	return BoundNodeKindConditionalGotoStatement
}
func (b *ConditionalGotoStatement) Type() reflect.Kind {
	return reflect.Invalid
}
