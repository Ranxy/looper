package syntax

type ElseClauseSyntax struct {
	ElseKeywords  SyntaxToken
	ElseStatement Statement
}

func NewElseClause(elseKeywords SyntaxToken, elsStatement Statement) *ElseClauseSyntax {
	return &ElseClauseSyntax{
		ElseKeywords:  elseKeywords,
		ElseStatement: elsStatement,
	}
}

func (e *ElseClauseSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.ElseKeywords, e.ElseStatement}
}

func (e *ElseClauseSyntax) Kind() SyntaxKind {
	return SyntaxKindElseKeywords
}
