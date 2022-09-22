package symbol

type LocalVariableSymbol struct {
	variableSymbol
}

func NewLocalVariableSymbol(name string, readOnly bool, tp *TypeSymbol) *LocalVariableSymbol {
	return &LocalVariableSymbol{
		variableSymbol: *newVariableSymbol(name, readOnly, tp),
	}
}
func (b *LocalVariableSymbol) Kind() SymbolKind {
	return SymbolKindLocalVariable
}
