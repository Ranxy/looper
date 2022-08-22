package bind

import (
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/syntax"
)

type BoundVariableExpression struct {
	Variable *syntax.VariableSymbol
}

func NewBoundVariableExpression(variable *syntax.VariableSymbol) *BoundVariableExpression {
	return &BoundVariableExpression{
		Variable: variable,
	}
}

func (b *BoundVariableExpression) Kind() BoundNodeKind {
	return BoundNodeKindVariableExpress
}
func (b *BoundVariableExpression) Type() reflect.Kind {
	return b.Variable.Type
}
func (b *BoundVariableExpression) GetChildren() []BoundNode {
	return []BoundNode{}
}
func (b *BoundVariableExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Variable}
}
