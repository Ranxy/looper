package bind

import (
	"fmt"
	"reflect"
)

type BoundWhileStatements struct {
	Condition BoundExpression
	Body      Boundstatement
}

func NewBoundWhileStatements(condition BoundExpression, body Boundstatement) *BoundWhileStatements {
	return &BoundWhileStatements{
		Condition: condition,
		Body:      body,
	}
}

func (b *BoundWhileStatements) Kind() BoundNodeKind {
	return BoundNodeKindWhileStatement
}
func (b *BoundWhileStatements) Type() reflect.Kind {
	return reflect.Invalid
}
func (b *BoundWhileStatements) GetChildren() []BoundNode {
	return []BoundNode{b.Condition, b.Body}
}
func (b *BoundWhileStatements) GetProperties() []fmt.Stringer {
	return []fmt.Stringer{}
}
