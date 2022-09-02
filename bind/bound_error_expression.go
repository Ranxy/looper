package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundErrorExpression struct {
}

func NewBoundErrorExpression() *BoundErrorExpression {
	return &BoundErrorExpression{}
}

func (b *BoundErrorExpression) Kind() BoundNodeKind {
	return BoundNodeKindErrorExpress
}
func (b *BoundErrorExpression) Type() *symbol.TypeSymbol {
	return symbol.TypeError
}

func (b *BoundErrorExpression) GetChildren() []BoundNode {
	return []BoundNode{}

}

func (b *BoundErrorExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
