package syntax

type UnitExpress struct {
	Open  SyntaxToken
	Close SyntaxToken
}

func NewUnitExpress(open SyntaxToken, close SyntaxToken) *UnitExpress {
	return &UnitExpress{
		Open:  open,
		Close: close,
	}
}

func (e *UnitExpress) GetChildren() []SyntaxNode {
	return []SyntaxNode{}
}

func (e *UnitExpress) Kind() SyntaxKind {
	return SyntaxKindUnitExpress
}
func (e *UnitExpress) String() string {
	return "()"
}
