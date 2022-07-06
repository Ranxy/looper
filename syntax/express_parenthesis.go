package syntax

import "fmt"

type ParenthesisExpress struct {
	kind  SyntaxKind
	Open  SyntaxToken
	Close SyntaxToken
	Expr  Express
}

func NewParenthesisExpress(open SyntaxToken, expr Express, close SyntaxToken) *ParenthesisExpress {

	return &ParenthesisExpress{
		kind:  SyntaxKindParenthesizedExpress,
		Open:  open,
		Close: close,
		Expr:  expr,
	}
}

func (e *ParenthesisExpress) GetChildren() []Express {
	return []Express{e.Open, e.Expr, e.Close}
}

func (e *ParenthesisExpress) Kind() SyntaxKind {
	return e.kind
}
func (e *ParenthesisExpress) String() string {
	return fmt.Sprintf("ParenthesisExpress:  Expr %s", e.Expr)
}
