package syntax

var (
	_ Express = &LiteralExpress{}
	_ Express = &BinaryExpress{}
	_ Express = &UnaryExpress{}
	_ Express = &ParenthesisExpress{}
)

type Express interface {
	Kind() SyntaxKind
	GetChildren() []Express
}
