package bind

import (
	"fmt"
	"reflect"

	"github.com/Ranxy/looper/syntax"
)

type BoundAssignmentExpression struct {
	Variable *syntax.VariableSymbol
	Express  BoundExpression
}

func NewBoundAssignmentExpression(variable *syntax.VariableSymbol, express BoundExpression) *BoundAssignmentExpression {
	return &BoundAssignmentExpression{
		Variable: variable,
		Express:  express,
	}
}

func (b *BoundAssignmentExpression) Kind() BoundNodeKind {
	return BoundNodeKindAssignmentExpress
}
func (b *BoundAssignmentExpression) Type() reflect.Kind {
	return b.Express.Type()
}

func (b *BoundAssignmentExpression) GetChildren() []BoundNode {
	return []BoundNode{b.Express}
}

func (b *BoundAssignmentExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Variable}
}
