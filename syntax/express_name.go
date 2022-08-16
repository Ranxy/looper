package syntax

type NameExpress struct {
	Identifier SyntaxToken
}

func NewNameExpress(Identifier SyntaxToken) *NameExpress {

	return &NameExpress{
		Identifier: Identifier,
	}
}

func (e *NameExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Identifier}
}
func (e *NameExpress) Kind() SyntaxKind {
	return SyntaxKindNameExpress
}
