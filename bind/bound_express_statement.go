package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
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
func (b *BoundExpressStatements) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}

func (b *BoundExpressStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Express}

}

func (b *BoundExpressStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
