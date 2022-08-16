package syntax

type SyntaxNode interface {
	GetChildren() []SyntaxNode
	Kind() SyntaxKind
}
