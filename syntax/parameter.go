package syntax

type ParameterSyntax struct {
	Identifier SyntaxToken
	Type       *TypeClauseSyntax
}

func NewParameterSyntax(identifier SyntaxToken, tp *TypeClauseSyntax) *ParameterSyntax {
	return &ParameterSyntax{
		Identifier: identifier,
		Type:       tp,
	}
}

func (e *ParameterSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Identifier, e.Type}
}
func (e *ParameterSyntax) Kind() SyntaxKind {
	return SyntaxKindParameter
}
