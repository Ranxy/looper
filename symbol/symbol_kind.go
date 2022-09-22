package symbol

type SymbolKind int

const (
	SymbolKindLocalVariable SymbolKind = iota
	SymbolKindGlobalVariable
	SymbolKindType
	SymbolKindFunction
	SymbolKindParameter
)
