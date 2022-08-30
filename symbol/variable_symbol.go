package symbol

type VariableSymbol struct {
	Name       string
	IsReadOnly bool
	Type       *TypeSymbol
}

func NewVariableSymbol(name string, readOnly bool, tp *TypeSymbol) *VariableSymbol {
	return &VariableSymbol{
		Name:       name,
		IsReadOnly: readOnly,
		Type:       tp,
	}
}
func (b *VariableSymbol) Kind() SymbolKind {
	return SymbolKindVariable
}

func (b *VariableSymbol) GetName() string {
	return b.Name
}

func (v *VariableSymbol) String() string {
	return v.Name + ":" + v.Type.String()
}
