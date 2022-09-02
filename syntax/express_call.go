package syntax

type CallExpress struct {
	Identifier       SyntaxToken
	OpenParenthesis  SyntaxToken
	Params           SeparatedList
	CloseParenthesis SyntaxToken
}

func NewCallExpress(identifier SyntaxToken, openParenthesis SyntaxToken, params SeparatedList, closeParenthises SyntaxToken) *CallExpress {
	return &CallExpress{
		Identifier:       identifier,
		OpenParenthesis:  openParenthesis,
		Params:           params,
		CloseParenthesis: closeParenthises,
	}
}

func (c *CallExpress) Kind() SyntaxKind {
	return SyntaxKindCallExpress
}

func (c *CallExpress) GetChildren() []SyntaxNode {
	res := []SyntaxNode{c.Identifier, c.OpenParenthesis}

	res = append(res, c.Params.List()...)
	res = append(res, c.CloseParenthesis)
	return res
}
