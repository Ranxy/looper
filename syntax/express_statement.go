package syntax

type ExpressStatement struct {
	Express Express
}

func NewExpressStatement(express Express) *ExpressStatement {

	return &ExpressStatement{
		Express: express,
	}
}

func (e *ExpressStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Express}
}
func (e *ExpressStatement) Kind() SyntaxKind {
	return SyntaxKindExpressStatement
}
