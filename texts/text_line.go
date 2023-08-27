package texts

type TextLine struct {
	Source                   *TextSource
	Start                    int
	Length                   int
	LengthIncludingLineBreak int
}

func NewTextLine(start, length, LengthIncludingLineBreak int, source *TextSource) TextLine {
	return TextLine{
		Source:                   source,
		Start:                    start,
		Length:                   length,
		LengthIncludingLineBreak: length,
	}
}

func (t *TextLine) End() int {
	return t.Start + t.Length
}

func (t *TextLine) Span() TextSpan {
	return NewTextSpan(t.Start, t.Length)
}

func (t *TextLine) SpanWithLineBreak() TextSpan {
	return NewTextSpan(t.Start, t.LengthIncludingLineBreak)
}

func (t *TextLine) String() string {
	return string(t.Source.Text[t.Start : t.Start+t.Length])
}

func (t *TextLine) TabCount(length int) int {
	cnt := 0
	for i := t.Start; i < t.Start+length; i++ {
		if t.Source.Text[i] == '\t' {
			cnt += 1
		}
	}
	return cnt
}
