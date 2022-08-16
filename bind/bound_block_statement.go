package bind

import (
	"reflect"
)

type BoundBlockStatements struct {
	Statement []Boundstatement
}

func NewBoundBlockStatement(statements []Boundstatement) *BoundBlockStatements {
	return &BoundBlockStatements{
		Statement: statements,
	}
}

func (b *BoundBlockStatements) Kind() BoundNodeKind {
	return BoundNodeKindBlockStatement
}
func (b *BoundBlockStatements) Type() reflect.Kind {
	return reflect.Invalid
}
