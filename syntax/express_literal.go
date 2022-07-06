package syntax

type LiteralExpress struct {
	kind    SyntaxKind
	Literal SyntaxToken
}

func NewLiteralExpress(literal SyntaxToken) *LiteralExpress {

	return &LiteralExpress{
		kind:    SyntaxKindBinaryExpress,
		Literal: literal,
	}
}

func (e *LiteralExpress) GetChildren() []Express {
	return []Express{e.Literal}
}
func (e *LiteralExpress) Kind() SyntaxKind {
	return e.kind
}
