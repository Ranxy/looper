package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundForStatements struct {
	InitCondition            Boundstatement
	EndCheckConditionExpress BoundExpression //must be bool
	UpdateCondition          Boundstatement
	Body                     Boundstatement
}

func NewBoundForStatements(initCondition Boundstatement, endCondition BoundExpression, updateCondition Boundstatement, body Boundstatement) *BoundForStatements {
	return &BoundForStatements{
		InitCondition:            initCondition,
		EndCheckConditionExpress: endCondition,
		UpdateCondition:          updateCondition,
		Body:                     body,
	}
}

func (b *BoundForStatements) Kind() BoundNodeKind {
	return BoundNodeKindForStatement
}
func (b *BoundForStatements) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}

func (b *BoundForStatements) GetChildren() []BoundNode {
	return []BoundNode{b.InitCondition, b.UpdateCondition, b.EndCheckConditionExpress, b.Body}

}

func (b *BoundForStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
