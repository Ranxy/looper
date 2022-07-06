package syntax

type UnaryExpress struct {
	kind     SyntaxKind
	Operator SyntaxToken
	Operand  Express
}

func NewUnaryExpress(operator SyntaxToken, operand Express) *UnaryExpress {

	return &UnaryExpress{
		kind:     SyntaxKindBinaryExpress,
		Operator: operator,
		Operand:  operand,
	}
}

func (e *UnaryExpress) GetChildren() []Express {
	return []Express{e.Operator, e.Operand}
}
func (e *UnaryExpress) Kind() SyntaxKind {
	return e.kind
}
