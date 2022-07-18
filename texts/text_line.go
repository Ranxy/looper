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
