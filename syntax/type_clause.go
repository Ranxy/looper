package syntax

type TypeClauseSyntax struct {
	ColonToken SyntaxToken
	Identifier SyntaxToken
}

func NewTypeClauseSyntax(colonToken, Identifier SyntaxToken) *TypeClauseSyntax {

	return &TypeClauseSyntax{
		ColonToken: colonToken,
		Identifier: Identifier,
	}
}

func (e *TypeClauseSyntax) GetChildren() []SyntaxNode {
	if e == nil {
		return []SyntaxNode{}
	}
	return []SyntaxNode{e.ColonToken, e.Identifier}
}
func (e *TypeClauseSyntax) Kind() SyntaxKind {
	return SyntaxKindTypeClause
}
