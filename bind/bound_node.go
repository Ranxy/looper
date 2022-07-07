package bind

type BoundNodeKind int

const (
	BoundNodeKindLiteralExpress BoundNodeKind = iota
	BoundNodeKindUnaryExpress
	BoundNodeKindBinaryExpress
)
