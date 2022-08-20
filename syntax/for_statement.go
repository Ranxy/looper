package syntax

type ForStatement struct {
	ForKeywords     SyntaxToken
	InitCondition   Express
	FirstSemicolon  SyntaxToken
	EndCondition    Express
	SecondSemicolon SyntaxToken
	UpdateCondition Express
	Body            Statement
}

func NewForStatement(forKeywords SyntaxToken,
	initCondition Statement,
	firstSemicolon SyntaxToken,
	endCondtion Express,
	secondSemicolon SyntaxToken,
	UpdateCondition Statement,
	body Statement) *ForStatement {
	return &ForStatement{
		ForKeywords:     forKeywords,
		InitCondition:   initCondition,
		FirstSemicolon:  firstSemicolon,
		EndCondition:    endCondtion,
		SecondSemicolon: secondSemicolon,
		UpdateCondition: UpdateCondition,
		Body:            body,
	}
}

func (e *ForStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.ForKeywords, e.InitCondition, e.FirstSemicolon,
		e.EndCondition, e.SecondSemicolon, e.UpdateCondition, e.Body}
}

func (s *ForStatement) Kind() SyntaxKind {
	return SyntaxkindForStatement
}
