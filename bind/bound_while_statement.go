package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundWhileStatements struct {
	BoundLoopStatements
	Condition BoundExpression
	Body      Boundstatement
}

func NewBoundWhileStatements(condition BoundExpression, body Boundstatement, breakLabel *BoundLabel, continueLabel *BoundLabel) *BoundWhileStatements {
	return &BoundWhileStatements{
		BoundLoopStatements: BoundLoopStatements{
			BreakLabel:    breakLabel,
			ContinueLabel: continueLabel,
		},
		Condition: condition,
		Body:      body,
	}
}

func (b *BoundWhileStatements) Kind() BoundNodeKind {
	return BoundNodeKindWhileStatement
}
func (b *BoundWhileStatements) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}
func (b *BoundWhileStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Condition, b.Body}
}
func (b *BoundWhileStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
