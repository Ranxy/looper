package bind

import (
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/syntax"
)

type BoundVariableDeclaration struct {
	Variable    *syntax.VariableSymbol
	Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable *syntax.VariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
	return &BoundVariableDeclaration{
		Variable:    variable,
		Initializer: initializer,
	}
}

func (b *BoundVariableDeclaration) Type() reflect.Kind {
	return reflect.Invalid
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
