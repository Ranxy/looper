package syntax

type CompliationUnit struct {
	Statements []MemberSyntax
	EofToken   SyntaxToken
}

func NewCompliationUnit(Statements []MemberSyntax, eofToken SyntaxToken) *CompliationUnit {
	return &CompliationUnit{
		Statements: Statements,
		EofToken:   eofToken,
	}
}

func (e *CompliationUnit) GetChildren() []SyntaxNode {
	res := []SyntaxNode{}
	for _, s := range e.Statements {
		res = append(res, s)
	}
	res = append(res, e.EofToken)
	return res
}
func (e *CompliationUnit) Kind() SyntaxKind {
	return SyntaxKindCompilationUnit
}
