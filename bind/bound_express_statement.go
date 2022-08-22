package bind

import (
	"fmt"
	"reflect"
)

type BoundExpressStatements struct {
	Express BoundExpression
}

func NewBoundExpressStatements(express BoundExpression) *BoundExpressStatements {
	return &BoundExpressStatements{
		Express: express,
	}
}

func (b *BoundExpressStatements) Kind() BoundNodeKind {
	return BoundNodeKindExpressionStatement
}
func (b *BoundExpressStatements) Type() reflect.Kind {
	return reflect.Invalid
}

func (b *BoundExpressStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Express}

}

func (b *BoundExpressStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
