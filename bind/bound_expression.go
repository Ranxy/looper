package bind

import (
	"fmt"

	"github.com/Ranxy/looper/symbol"
)

type BoundNode interface {
	Type() *symbol.TypeSymbol
	Kind() BoundNodeKind
	NodePrint
}

type NodePrint interface {
	GetChildren() []BoundNode
	GetProperties() []fmt.Stringer
}

type BoundExpression interface {
	BoundNode
}

type Boundstatement interface {
	BoundNode
}
