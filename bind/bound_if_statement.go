package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundIfStatements struct {
	Condition     BoundExpression
	ThenStatement Boundstatement
	ElseStatement Boundstatement
}

func NewBoundIfStatements(condition BoundExpression, thenStatement, elseStatement Boundstatement) *BoundIfStatements {
	return &BoundIfStatements{
		Condition:     condition,
		ThenStatement: thenStatement,
		ElseStatement: elseStatement,
	}
}

func (b *BoundIfStatements) Kind() BoundNodeKind {
	return BoundNodeKindIfStatement
}
func (b *BoundIfStatements) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}

func (b *BoundIfStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Condition, b.ThenStatement, b.ElseStatement}

}

func (b *BoundIfStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
