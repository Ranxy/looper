package bind

import "reflect"

type BoundNode interface {
	Type() reflect.Kind
	Kind() BoundNodeKind
}

type BoundExpression interface {
	BoundNode
}

type Boundstatement interface {
	BoundNode
}
