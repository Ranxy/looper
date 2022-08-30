package bind

type BoundLabel struct {
	Name string
}

func NewBoundLabel(name string) *BoundLabel {
	return &BoundLabel{
		Name: name,
	}
}

func (l *BoundLabel) String() string {
	return l.Name
}
