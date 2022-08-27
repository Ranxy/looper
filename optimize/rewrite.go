package optimize

var _ Rewrite = &BasicRewrite{}

type Rewrite interface {
	RewriteStatement
	RewriteExpression
}
