package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type GotoStatement struct {
	Label *BoundLabel
}

func NewGotoSymbol(label *BoundLabel) *GotoStatement {
	return &GotoStatement{
		Label: label,
	}
}

func (b *GotoStatement) Kind() BoundNodeKind {
	return BoundNodeKindGotoStatement
}
func (b *GotoStatement) Type() *symbol.TypeSymbol {
	return symbol.TypeUnkonw
}

func (b *GotoStatement) GetChildren() []BoundNode {
	return []BoundNode{}
}
func (b *GotoStatement) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Label}
}
