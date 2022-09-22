package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundVariableExpression struct {
	Variable symbol.VariableSymbol
}

func NewBoundVariableExpression(variable symbol.VariableSymbol) *BoundVariableExpression {
	return &BoundVariableExpression{
		Variable: variable,
	}
}

func (b *BoundVariableExpression) Kind() BoundNodeKind {
	return BoundNodeKindVariableExpress
}
func (b *BoundVariableExpression) Type() *symbol.TypeSymbol {
	return b.Variable.GetType()
}
func (b *BoundVariableExpression) GetChildren() []BoundNode {
	return []BoundNode{}
}
func (b *BoundVariableExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Variable}
}
