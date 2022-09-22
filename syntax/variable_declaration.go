package syntax

type VariableDeclarationSyntax struct {
	Keyword     SyntaxToken
	Identifier  SyntaxToken
	TypeClause  *TypeClauseSyntax
	Equals      SyntaxToken
	Initializer Express
}

func NewVariableDeclarationSyntax(keyword, Identifier SyntaxToken, typeClause *TypeClauseSyntax, equal SyntaxToken, initializer Express) *VariableDeclarationSyntax {

	return &VariableDeclarationSyntax{
		Keyword:     keyword,
		Identifier:  Identifier,
		TypeClause:  typeClause,
		Equals:      equal,
		Initializer: initializer,
	}
}

func (e *VariableDeclarationSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Keyword, e.Identifier, e.TypeClause, e.Equals, e.Initializer}
}
func (e *VariableDeclarationSyntax) Kind() SyntaxKind {
	return SyntaxKindVariableDeclaration
}
