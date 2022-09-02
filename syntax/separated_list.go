package syntax

type SeparatedList struct {
	nodeAndSeparators []SyntaxNode
}

func NewSeptartedList(list []SyntaxNode) *SeparatedList {
	return &SeparatedList{
		nodeAndSeparators: list,
	}
}

func (s *SeparatedList) Count() int {
	return (len(s.nodeAndSeparators) + 1) / 2
}

func (s *SeparatedList) Get(x int) SyntaxNode {
	return s.nodeAndSeparators[x*2]
}
func (s *SeparatedList) List() []SyntaxNode {
	count := s.Count()
	res := make([]SyntaxNode, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, s.Get(i))
	}
	return res
}
func (s *SeparatedList) ListWithSpearated() []SyntaxNode {
	return s.nodeAndSeparators
}

func (s *SeparatedList) GetSeparator(i int) *SyntaxToken {
	if i == s.Count()-1 {
		return nil
	}
	v := any(&s.nodeAndSeparators[i*2+1])
	return v.(*SyntaxToken)
}
