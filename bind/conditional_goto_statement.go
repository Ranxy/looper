package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type ConditionalGotoStatement struct {
	Label       *BoundLabel
	Condition   BoundExpression
	JumpIfFalse bool
}

func NewConditionalGotoSymbol(label *BoundLabel, condition BoundExpression, jumpIfFalse bool) *ConditionalGotoStatement {
	return &ConditionalGotoStatement{
		Label:       label,
		Condition:   condition,
		JumpIfFalse: jumpIfFalse,
	}
}

func (b *ConditionalGotoStatement) Kind() BoundNodeKind {
	return BoundNodeKindConditionalGotoStatement
}
func (b *ConditionalGotoStatement) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}
func (b *ConditionalGotoStatement) GetChildren() []BoundNode {
	return []BoundNode{b.Condition}
}
func (b *ConditionalGotoStatement) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Label, newBookStringer(b.JumpIfFalse)}
}
