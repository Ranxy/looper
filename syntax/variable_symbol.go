package syntax

import "reflect"

type VariableSymbol struct {
	Name string
	Type reflect.Kind
}

func NewVariableSymbol(name string, tp reflect.Kind) *VariableSymbol {
	return &VariableSymbol{
		Name: name,
		Type: tp,
	}
}
