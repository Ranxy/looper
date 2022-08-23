package bind

import (
	"fmt"
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

func (b *BoundBlockStatements) GetChildren() []BoundNode {
	res := []BoundNode{}
	for _, s := range b.Statement {
		res = append(res, s)
	}
	return res
}

func (b *BoundBlockStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
