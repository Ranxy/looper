package bind

type BoundNodeKind int

const (
	BoundNodeKindLiteralExpress BoundNodeKind = iota
	BoundNodeKindVariableExpress
	BoundNodeKindAssignmentExpress
	BoundNodeKindUnaryExpress
	BoundNodeKindBinaryExpress

	BoundNodeKindBlockStatement
	BoundNodeKindVariableDeclaration
	BoundNodeKindIfStatement
	BoundNodeKindExpressionStatement
)
