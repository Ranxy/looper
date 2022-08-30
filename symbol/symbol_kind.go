package symbol

type SymbolKind int

const (
	SymbolKindVariable SymbolKind = iota
	SymbolKindType
)
