package syntax

type UnaryExpress struct {
	Operator SyntaxToken
	Operand  Express
}

func NewUnaryExpress(operator SyntaxToken, operand Express) *UnaryExpress {

	return &UnaryExpress{
		Operator: operator,
		Operand:  operand,
	}
}

func (e *UnaryExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Operator, e.Operand}
}
func (e *UnaryExpress) Kind() SyntaxKind {
	return SyntaxKindUnaryExpress
}
