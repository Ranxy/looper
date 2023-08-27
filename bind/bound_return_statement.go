package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundReturnStatements struct {
	Express BoundExpression
}

func NewBoundReturnStatements(express BoundExpression) *BoundReturnStatements {
	return &BoundReturnStatements{
		Express: express,
	}
}

func (b *BoundReturnStatements) Kind() BoundNodeKind {
	return BoundNodeKindReturnStatement
}

func (b *BoundReturnStatements) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}

func (b *BoundReturnStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Express}

}

func (b *BoundReturnStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
