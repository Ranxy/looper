package symbol

type Symbol interface {
	Kind() SymbolKind
	GetName() string
	String() string
}
