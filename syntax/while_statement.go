package syntax

type WhileStatement struct {
	WhileKeywords SyntaxToken
	Condition     Express
	Body          Statement
}

func NewWhileStatement(whileKeywords SyntaxToken, condition Express, statement Statement) *WhileStatement {
	return &WhileStatement{
		WhileKeywords: whileKeywords,
		Condition:     condition,
		Body:          statement,
	}
}

func (e *WhileStatement) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.WhileKeywords, e.Condition, e.Body}
}

func (s *WhileStatement) Kind() SyntaxKind {
	return SyntaxKindWhileStatement
}
