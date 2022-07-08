package bind

import (
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
