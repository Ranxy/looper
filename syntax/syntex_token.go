package syntax

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
