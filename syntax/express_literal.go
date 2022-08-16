package syntax

type LiteralExpress struct {
	Literal SyntaxToken
	Value   any
}

func NewLiteralExpress(literal SyntaxToken) *LiteralExpress {

	return &LiteralExpress{
		Literal: literal,
		Value:   literal.Value,
	}
}
func NewLiteralValueExpress(literal SyntaxToken, value any) *LiteralExpress {

	return &LiteralExpress{
		Literal: literal,
		Value:   value,
	}
}
func (e *LiteralExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Literal}
}
func (e *LiteralExpress) Kind() SyntaxKind {
	return SyntaxKindLiteralExpress
}
