package symbol

type VariableSymbol interface {
	GetName() string
	String() string
	Kind() SymbolKind
	GetType() *TypeSymbol
	IsReadOnly() bool
}
type variableSymbol struct {
	Name       string
	isReadOnly bool
	Type       *TypeSymbol
}

func newVariableSymbol(name string, readOnly bool, tp *TypeSymbol) *variableSymbol {
	return &variableSymbol{
		Name:       name,
		isReadOnly: readOnly,
		Type:       tp,
	}
}

func (b *variableSymbol) GetName() string {
	return b.Name
}

func (v *variableSymbol) String() string {
	return v.Name + ":" + v.Type.String()
}
func (v *variableSymbol) GetType() *TypeSymbol {
	return v.Type
}

func (v *variableSymbol) IsReadOnly() bool {
	return v.isReadOnly
}
