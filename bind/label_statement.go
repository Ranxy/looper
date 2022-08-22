package bind

import (
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/label"
)

type LabelStatement struct {
	Label *label.LabelSymbol
}

func NewLabelSymbol(label *label.LabelSymbol) *LabelStatement {
	return &LabelStatement{
		Label: label,
	}
}

func (b *LabelStatement) Kind() BoundNodeKind {
	return BoundNodeKindLabelStatement
}
func (b *LabelStatement) Type() reflect.Kind {
	return reflect.Invalid
}
func (b *LabelStatement) GetChildren() []BoundNode {
	return []BoundNode{}
}
func (b *LabelStatement) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Label}
}
