package bind

import "reflect"

type BoundExpression interface {
	Type() reflect.Kind
	Kind() BoundNodeKind
}
