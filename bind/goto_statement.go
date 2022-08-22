package bind

import (
	"reflect"

	"github.com/Ranxy/looper/label"
)

type GotoStatement struct {
	Label *label.LabelSymbol
}

func NewGotoSymbol(label *label.LabelSymbol) *GotoStatement {
	return &GotoStatement{
		Label: label,
	}
}

func (b *GotoStatement) Kind() BoundNodeKind {
	return BoundNodeKindGotoStatement
}
func (b *GotoStatement) Type() reflect.Kind {
	return reflect.Invalid
}
