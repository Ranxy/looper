package syntax

type BinaryExpress struct {
	kind     SyntaxKind
	Operator SyntaxToken
	Left     Express
	Right    Express
}

func NewBinaryExpress(left Express, operator SyntaxToken, right Express) *BinaryExpress {

	return &BinaryExpress{
		kind:     SyntaxKindBinaryExpress,
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (e *BinaryExpress) GetChildren() []Express {
	return []Express{e.Left, e.Operator, e.Right}
}

func (e *BinaryExpress) Kind() SyntaxKind {
	return e.kind
}
