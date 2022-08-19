package syntax

type IfStatement struct {
	IfKeywords    SyntaxToken
	Condition     Express
	ThenStatement Statement
	ElseClause    *ElseClauseSyntax
}

func NewIfStatement(ifKeywords SyntaxToken, condition Express, thenStatement Statement, elseClause *ElseClauseSyntax) *IfStatement {
	return &IfStatement{
		IfKeywords:    ifKeywords,
		Condition:     condition,
		ThenStatement: thenStatement,
		ElseClause:    elseClause,
	}
}

func (e *IfStatement) GetChildren() []SyntaxNode {
	res := []SyntaxNode{e.IfKeywords, e.Condition, e.ThenStatement}
	if e.ElseClause != nil {
		res = append(res, e.ElseClause)
	}

	return res
}

func (s *IfStatement) Kind() SyntaxKind {
	return SyntaxKindIfStatement
}
