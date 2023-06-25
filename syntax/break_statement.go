package syntax

type BreakStatement struct {
	Keywords SyntaxToken
}

func NewBreakStatement(keywords SyntaxToken) *BreakStatement {

	return &BreakStatement{
		Keywords: keywords,
	}
}

func (e *BreakStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Keywords}
}
func (e *BreakStatement) Kind() SyntaxKind {
	return SyntaxKindBreakStatement
}
