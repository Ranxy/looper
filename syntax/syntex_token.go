package syntax

import "github.com/Ranxy/looper/diagnostic"

type SyntaxToken struct {
	kind     SyntaxKind
	Position int
	Text     string
	Value    any
}

func (s SyntaxToken) GetChildren() []Express {
	return []Express{}
}
func (e SyntaxToken) Kind() SyntaxKind {
	return e.kind
}

func (e SyntaxToken) Span() diagnostic.TextSpan {
	return diagnostic.NewTextSpan(e.Position, len([]rune(e.Text)))
}
