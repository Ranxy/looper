package bind

import (
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
