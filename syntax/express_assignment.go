package syntax

type AssignmentExpress struct {
	kind       SyntaxKind
	Identifier SyntaxToken
	Equal      SyntaxToken
	Express    Express
}

func NewAssignmentExpress(identifier SyntaxToken, equal SyntaxToken, express Express) *AssignmentExpress {
	return &AssignmentExpress{
		kind:       SyntaxKindAssignmentExpress,
		Identifier: identifier,
		Equal:      equal,
		Express:    express,
	}
}

func (e *AssignmentExpress) GetChildren() []Express {
	return []Express{e.Identifier, e.Equal, e.Express}
}
func (e *AssignmentExpress) Kind() SyntaxKind {
	return e.kind
}
