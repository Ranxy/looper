package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type LabelStatement struct {
	Label *BoundLabel
}

func NewLabelSymbol(label *BoundLabel) *LabelStatement {
	return &LabelStatement{
		Label: label,
	}
}

func (b *LabelStatement) Kind() BoundNodeKind {
	return BoundNodeKindLabelStatement
}
func (b *LabelStatement) Type() *symbol.TypeSymbol {
	return symbol.TypeUnkonw
}
func (b *LabelStatement) GetChildren() []BoundNode {
	return []BoundNode{}
}
func (b *LabelStatement) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Label}
}
