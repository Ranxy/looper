package syntax

type BinaryExpress struct {
	Operator SyntaxToken
	Left     Express
	Right    Express
}

func NewBinaryExpress(left Express, operator SyntaxToken, right Express) *BinaryExpress {

	return &BinaryExpress{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (e *BinaryExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Left, e.Operator, e.Right}
}

func (e *BinaryExpress) Kind() SyntaxKind {
	return SyntaxKindBinaryExpress
}
