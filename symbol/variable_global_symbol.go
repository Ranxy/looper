package symbol

type GlobalVariableSymbol struct {
	variableSymbol
}

func NewGlobalVariableSymbol(name string, readOnly bool, tp *TypeSymbol) *GlobalVariableSymbol {
	return &GlobalVariableSymbol{
		variableSymbol: *newVariableSymbol(name, readOnly, tp),
	}
}
func (b *GlobalVariableSymbol) Kind() SymbolKind {
	return SymbolKindGlobalVariable
}
