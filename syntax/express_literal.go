package syntax

type LiteralExpress struct {
	kind    SyntaxKind
	Literal SyntaxToken
	Value   any
}

func NewLiteralExpress(literal SyntaxToken) *LiteralExpress {

	return &LiteralExpress{
		kind:    SyntaxKindLiteralExpress,
		Literal: literal,
		Value:   literal.Value,
	}
}
func NewLiteralValueExpress(literal SyntaxToken, value any) *LiteralExpress {

	return &LiteralExpress{
		kind:    SyntaxKindLiteralExpress,
		Literal: literal,
		Value:   value,
	}
}
func (e *LiteralExpress) GetChildren() []Express {
	return []Express{e.Literal}
}
func (e *LiteralExpress) Kind() SyntaxKind {
	return e.kind
}
