package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

type BoundUnitExpression struct {
	Unit *syntax.UnitExpress
}

func NewBoundUnitExpression(unit *syntax.UnitExpress) *BoundUnitExpression {
	return &BoundUnitExpression{
		Unit: unit,
	}
}
func (b *BoundUnitExpression) Type() *symbol.TypeSymbol {
	return symbol.TypeUnit
}
func (b *BoundUnitExpression) Kind() BoundNodeKind {
	return BoundNodeKindUnitExpress
}

func (b *BoundUnitExpression) GetChildren() []BoundNode {
	return []BoundNode{}

}
func (b *BoundUnitExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Unit}
}
