package symbol

type TypeSymbol struct {
	name string
}

func NewTypeSymbol(name string) *TypeSymbol {
	return &TypeSymbol{name: name}
}

func (b *TypeSymbol) Kind() SymbolKind {
	return SymbolKindType
}

func (b *TypeSymbol) GetName() string {
	return b.name
}

func (v *TypeSymbol) String() string {
	return v.name
}

var (
	TypeError  = NewTypeSymbol("?")
	TypeString = NewTypeSymbol("string")
	TypeBool   = NewTypeSymbol("bool")
	TypeInt    = NewTypeSymbol("int")
	TypeUnit   = NewTypeSymbol("()")
)
