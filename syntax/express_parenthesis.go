package syntax

import "fmt"

type ParenthesisExpress struct {
	Open  SyntaxToken
	Close SyntaxToken
	Expr  Express
}

func NewParenthesisExpress(open SyntaxToken, expr Express, close SyntaxToken) *ParenthesisExpress {

	return &ParenthesisExpress{
		Open:  open,
		Close: close,
		Expr:  expr,
	}
}

func (e *ParenthesisExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{e.Open, e.Expr, e.Close}
}

func (e *ParenthesisExpress) Kind() SyntaxKind {
	return SyntaxKindParenthesizedExpress
}
func (e *ParenthesisExpress) String() string {
	return fmt.Sprintf("ParenthesisExpress:  Expr %s", e.Expr)
}
