package syntax

type CompliationUnit struct {
	Statement Statement
	EofToken  SyntaxToken
}

func NewCompliationUnit(Statement Statement, eofToken SyntaxToken) *CompliationUnit {
	return &CompliationUnit{
		Statement: Statement,
		EofToken:  eofToken,
	}
}

func (e *CompliationUnit) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Statement, e.EofToken}
}
func (e *CompliationUnit) Kind() SyntaxKind {
	return SyntaxKindCompilationUnit
}
