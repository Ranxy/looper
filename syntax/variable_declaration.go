package syntax

type VariableDeclarationSyntax struct {
	Keyword     SyntaxToken
	Identifier  SyntaxToken
	Equals      SyntaxToken
	Initializer Express
}

func NewVariableDeclarationSyntax(keyword, Identifier, equal SyntaxToken, initializer Express) *VariableDeclarationSyntax {

	return &VariableDeclarationSyntax{
		Keyword:     keyword,
		Identifier:  Identifier,
		Equals:      equal,
		Initializer: initializer,
	}
}

func (e *VariableDeclarationSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Keyword, e.Identifier, e.Equals, e.Initializer}
}
func (e *VariableDeclarationSyntax) Kind() SyntaxKind {
	return SyntaxKindVariableDeclaration
}
