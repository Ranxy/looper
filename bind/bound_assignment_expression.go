package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundAssignmentExpression struct {
	Variable *symbol.VariableSymbol
	Express  BoundExpression
}

func NewBoundAssignmentExpression(variable *symbol.VariableSymbol, express BoundExpression) *BoundAssignmentExpression {
	return &BoundAssignmentExpression{
		Variable: variable,
		Express:  express,
	}
}

func (b *BoundAssignmentExpression) Kind() BoundNodeKind {
	return BoundNodeKindAssignmentExpress
}
func (b *BoundAssignmentExpression) Type() *symbol.TypeSymbol {
	return b.Express.Type()
}

func (b *BoundAssignmentExpression) GetChildren() []BoundNode {
	return []BoundNode{b.Express}
}

func (b *BoundAssignmentExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Variable}
}
