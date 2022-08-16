package syntax

type AssignmentExpress struct {
	Identifier SyntaxToken
	Equal      SyntaxToken
	Express    Express
}

func NewAssignmentExpress(identifier SyntaxToken, equal SyntaxToken, express Express) *AssignmentExpress {
	return &AssignmentExpress{
		Identifier: identifier,
		Equal:      equal,
		Express:    express,
	}
}

func (e *AssignmentExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Identifier, e.Equal, e.Express}
}
func (e *AssignmentExpress) Kind() SyntaxKind {
	return SyntaxKindAssignmentExpress
}
