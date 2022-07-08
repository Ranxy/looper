package syntax

type NameExpress struct {
	kind       SyntaxKind
	Identifier SyntaxToken
}

func NewNameExpress(Identifier SyntaxToken) *NameExpress {

	return &NameExpress{
		kind:       SyntaxKindNameExpress,
		Identifier: Identifier,
	}
}

func (e *NameExpress) GetChildren() []Express {
	return []Express{e.Identifier}
}
func (e *NameExpress) Kind() SyntaxKind {
	return e.kind
}
