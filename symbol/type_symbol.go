package symbol

type TypeSymbol struct {
	name string
}

func NewTypeSymbol(name string) *TypeSymbol {
	return &TypeSymbol{name: name}
}

func (b *TypeSymbol) Kind() SymbolKind {
	return SymbolKindVariable
}

func (b *TypeSymbol) GetName() string {
	return b.name
}

func (v *TypeSymbol) String() string {
	return "type:" + v.name
}

var (
	TypeString = NewTypeSymbol("string")
	TypeBool   = NewTypeSymbol("bool")
	TypeInt    = NewTypeSymbol("int")
	TypeUnkonw = NewTypeSymbol("?")
)
