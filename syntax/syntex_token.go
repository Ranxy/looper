package syntax

import (
	"github.com/Ranxy/looper/texts"
)

type SyntaxToken struct {
	kind     SyntaxKind
	Position int
	Text     string
	Value    any
}

func (s SyntaxToken) GetChildren() []SyntaxNode {
	return []SyntaxNode{}
}
func (e SyntaxToken) Kind() SyntaxKind {
	return e.kind
}

func (e SyntaxToken) Span() texts.TextSpan {
	return texts.NewTextSpan(e.Position, len([]rune(e.Text)))
}
