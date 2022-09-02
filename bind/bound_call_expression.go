package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundCallExpression struct {
	Function  *symbol.FunctionSymbol
	Arguments []BoundExpression
}

func NewBoundcallExpression(function *symbol.FunctionSymbol, arguments []BoundExpression) *BoundCallExpression {
	return &BoundCallExpression{
		Function:  function,
		Arguments: arguments,
	}
}

func (b *BoundCallExpression) Type() *symbol.TypeSymbol {
	return b.Function.Type
}

func (b *BoundCallExpression) Kind() BoundNodeKind {
	return BoundNodeKindCallExpress
}
func (b *BoundCallExpression) GetChildren() []BoundNode {
	res := []BoundNode{}
	for _, arg := range b.Arguments {
		res = append(res, arg)
	}
	return res
}

func (b *BoundCallExpression) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{b.Function}
}
