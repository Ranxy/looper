package syntax

type FunctionDeclarationSyntax struct {
	Keyword          SyntaxToken
	Identifier       SyntaxToken
	OpenParenthesis  SyntaxToken
	ParameterList    *SeparatedList
	CloseParenthesis SyntaxToken
	Type             *TypeClauseSyntax
	Body             *BlockStatement
}

func NewFunctionDeclaration(keyword, identifier, openParenthesis SyntaxToken,
	parameterList *SeparatedList,
	closeParenthesis SyntaxToken, tp *TypeClauseSyntax, body *BlockStatement) *FunctionDeclarationSyntax {

	return &FunctionDeclarationSyntax{
		Keyword:          keyword,
		Identifier:       identifier,
		OpenParenthesis:  openParenthesis,
		ParameterList:    parameterList,
		CloseParenthesis: closeParenthesis,
		Type:             tp,
		Body:             body,
	}
}
func (e *FunctionDeclarationSyntax) GetChildren() []SyntaxNode {
	res := []SyntaxNode{e.Keyword, e.Identifier, e.OpenParenthesis}
	res = append(res, e.ParameterList.List()...)
	res = append(res, e.CloseParenthesis, e.Type, e.Body)
	return res
}

func (s *FunctionDeclarationSyntax) Kind() SyntaxKind {
	return SyntaxKindFunctionDeclaration
}
