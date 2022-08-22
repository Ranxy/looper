package syntax

import "reflect"

type VariableSymbol struct {
	Name       string
	IsReadOnly bool
	Type       reflect.Kind
}

func NewVariableSymbol(name string, readOnly bool, tp reflect.Kind) *VariableSymbol {
	return &VariableSymbol{
		Name:       name,
		IsReadOnly: readOnly,
		Type:       tp,
	}
}

func (v *VariableSymbol) String() string {
	return v.Name + ":" + v.Type.String()
}
