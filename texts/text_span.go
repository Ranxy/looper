package texts

type TextSpan struct {
	start  int
	length int
}

func NewTextSpan(start, length int) TextSpan {
	return TextSpan{
		start:  start,
		length: length,
	}
}

func (t TextSpan) Start() int {
	return t.start
}

func (t TextSpan) Length() int {
	return t.length
}

func (t TextSpan) End() int {
	return t.start + t.length
}
