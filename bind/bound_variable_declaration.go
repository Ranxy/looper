package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundVariableDeclaration struct {
	Variable    *symbol.VariableSymbol
	Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable *symbol.VariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
	return &BoundVariableDeclaration{
		Variable:    variable,
		Initializer: initializer,
	}
}

func (b *BoundVariableDeclaration) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}
func (b *BoundVariableDeclaration) Kind() BoundNodeKind {
	return BoundNodeKindVariableDeclaration
}
func (b *BoundVariableDeclaration) GetChildren() []BoundNode {
	return []BoundNode{b.Initializer}
}
func (b *BoundVariableDeclaration) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Variable}
}
