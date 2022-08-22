package label

type LabelSymbol struct {
	Name string
}

func NewLabelSymbol(name string) *LabelSymbol {
	return &LabelSymbol{
		Name: name,
	}
}

func (l *LabelSymbol) String() string {
	return l.Name
}
