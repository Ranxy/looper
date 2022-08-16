package syntax

type BlockStatement struct {
	OpenBraceToken  SyntaxToken
	Statements      []Statement
	CloseBraceToken SyntaxToken
}

func NewBlockStatement(open SyntaxToken, Statement []Statement, close SyntaxToken) *BlockStatement {
	return &BlockStatement{
		OpenBraceToken:  open,
		Statements:      Statement,
		CloseBraceToken: close,
	}
}

func (e *BlockStatement) GetChildren() []SyntaxNode {
	res := make([]SyntaxNode, 0, len(e.Statements)+2)
	res = append(res, e.OpenBraceToken)
	for _, s := range e.Statements {
		res = append(res, s)
	}
	res = append(res, e.CloseBraceToken)
	return res
}
func (e *BlockStatement) Kind() SyntaxKind {
	return SyntaxKindBlockStatement
}
