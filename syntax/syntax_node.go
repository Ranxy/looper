package syntax

import (
	"github.com/Ranxy/looper/texts"
)

type MemberSyntax interface {
	SyntaxNode
}

type SyntaxNode interface {
	GetChildren() []SyntaxNode
	Kind() SyntaxKind
}

func SyntaxNodeSpan(n SyntaxNode) texts.TextSpan {
	if st, ok := n.(SyntaxToken); ok {
		return st.Span()
	}

	childen := n.GetChildren()
	if len(childen) == 0 {
		return texts.NewTextSpan(0, 0)
	}
	first := childen[0]
	if len(childen) == 1 {
		return SyntaxNodeSpan(first)
	}
	second := childen[len(childen)-1]

	start := SyntaxNodeSpan(first).Start()
	end := SyntaxNodeSpan(second).End()
	return texts.NewTextSpan(start, end-start)
}

func GetLastToken(n SyntaxNode) *SyntaxToken {
	if t, ok := n.(*SyntaxToken); ok {
		return t
	}

	children := n.GetChildren()

	return GetLastToken(children[len(children)-1])
}
