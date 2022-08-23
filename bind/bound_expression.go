package bind

import (
	"fmt"
	"reflect"
)

type BoundNode interface {
	Type() reflect.Kind
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
