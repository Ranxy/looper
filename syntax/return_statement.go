package syntax

type ReturnStatement struct {
	ReturnKeywords SyntaxToken
	Express        Express
}

func NewReturnStatement(returnKeywords SyntaxToken,
	express Express) *ReturnStatement {
	return &ReturnStatement{
		ReturnKeywords: returnKeywords,
		Express:        express,
	}
}

func (e *ReturnStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.ReturnKeywords, e.Express}
}

func (s *ReturnStatement) Kind() SyntaxKind {
	return SyntaxKindReturnStatement
}
