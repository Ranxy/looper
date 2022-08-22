package bind

import (
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
