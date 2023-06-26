package syntax

type ContinueStatement struct {
	Keywords SyntaxToken
}

func NewContinueStatement(keywords SyntaxToken) *ContinueStatement {

	return &ContinueStatement{
		Keywords: keywords,
	}
}

func (e *ContinueStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Keywords}
}
func (e *ContinueStatement) Kind() SyntaxKind {
	return SyntaxKindContinueStatement
}
