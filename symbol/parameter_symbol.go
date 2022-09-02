package symbol

type ParameterSymbol struct {
	name string
	Type *TypeSymbol
}

func NewParameterSymbol(name string, tp *TypeSymbol) *ParameterSymbol {
	return &ParameterSymbol{
		name: name,
		Type: tp,
	}
}

func (s *ParameterSymbol) Kind() SymbolKind {
	return SymbolKindParameter
}

func (s *ParameterSymbol) GetName() string {
	return s.name
}

func (s *ParameterSymbol) String() string {
	return s.name + ":" + s.Type.GetName()
}
