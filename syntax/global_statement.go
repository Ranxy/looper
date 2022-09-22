package syntax

type GlobalStatement struct {
	Statement Statement
}

func NewGlobalStatement(statement Statement) *GlobalStatement {
	return &GlobalStatement{
		Statement: statement,
	}
}

func (e *GlobalStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Statement}
}
func (e *GlobalStatement) Kind() SyntaxKind {
	return SyntaxKindGlobalStatement
}
