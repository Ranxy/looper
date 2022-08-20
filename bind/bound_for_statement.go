package bind

import "reflect"

type BoundForStatements struct {
	InitCondition   Boundstatement
	EndCondition    BoundExpression //must be bool
	UpdateCondition Boundstatement
	Body            Boundstatement
}

func NewBoundForStatements(initCondition Boundstatement, endCondition BoundExpression, updateCondition Boundstatement, body Boundstatement) *BoundForStatements {
	return &BoundForStatements{
		InitCondition:   initCondition,
		EndCondition:    endCondition,
		UpdateCondition: updateCondition,
		Body:            body,
	}
}

func (b *BoundForStatements) Kind() BoundNodeKind {
	return BoundNodeKindForStatement
}
func (b *BoundForStatements) Type() reflect.Kind {
	return reflect.Invalid
}
