package bind

import (
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/label"
)

type ConditionalGotoStatement struct {
	Label       *label.LabelSymbol
	Condition   BoundExpression
	JumpIfFalse bool
}

func NewConditionalGotoSymbol(label *label.LabelSymbol, condition BoundExpression, jumpIfFalse bool) *ConditionalGotoStatement {
	return &ConditionalGotoStatement{
		Label:       label,
		Condition:   condition,
		JumpIfFalse: jumpIfFalse,
	}
}

func (b *ConditionalGotoStatement) Kind() BoundNodeKind {
	return BoundNodeKindConditionalGotoStatement
}
func (b *ConditionalGotoStatement) Type() reflect.Kind {
	return reflect.Invalid
}
func (b *ConditionalGotoStatement) GetChildren() []BoundNode {
	return []BoundNode{b.Condition}
}
func (b *ConditionalGotoStatement) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Label, newBookStringer(b.JumpIfFalse)}
}
