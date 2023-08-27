package texts

import (
	"unicode"
)

type TextSource struct {
	Text  []rune
	Lines []TextLine
}

func NewTextSource(text []rune) *TextSource {

	s := &TextSource{
		Text: text,
	}
	s.ParserLines(text)
	return s
}

func (s *TextSource) Len() int {
	return len(s.Text)
}
func (s *TextSource) RuneAt(idx int) rune {
	return s.Text[idx]
}

func (s *TextSource) ParserLines(text []rune) {
	pos := 0
	lineStart := 0

	for pos < len(text) {
		lineBreakWidth := getLineBreakWidth(s.Text, pos)
		if lineBreakWidth == 0 {
			pos += 1
		} else {
			//add line
			s.addLine(pos, lineStart, lineBreakWidth)
			pos += lineBreakWidth
			lineStart = pos
		}
	}
	if pos >= lineStart {
		s.addLine(pos, lineStart, 0)
	}
}

func (s *TextSource) addLine(pos int, lineStart int, lineBreakWidth int) {
	lineLength := pos - lineStart
	lineLengthWithLineBreak := lineLength + lineBreakWidth

	s.Lines = append(s.Lines, NewTextLine(lineStart, lineLength, lineLengthWithLineBreak, s))
}

func getLineBreakWidth(text []rune, pos int) int {
	c := text[pos]
	var l rune
	if pos+1 >= len(text) {
		l = unicode.MaxRune
	} else {
		l = text[pos+1]
	}

	if c == '\r' && l == '\n' {
		return 2
	}
	if c == '\r' || c == '\n' {
		return 1
	}
	return 0
}

func (s *TextSource) String() string {
	return string(s.Text)
}

func (s *TextSource) StringSub(start, end int) string {
	return string(s.Text[start:end])
}
func (s *TextSource) StringSpan(span TextSpan) string {
	return s.StringSub(span.start, span.End())
}

func (s *TextSource) GetLineIndex(pos int) int {
	lower := 0
	upper := len(s.Lines) - 1

	for lower <= upper {
		idx := lower + (upper-lower)/2
		start := s.Lines[idx].Start

		if pos == start {
			return idx
		}
		if pos < start {
			upper = idx - 1
		} else {
			lower = idx + 1
		}
	}
	return lower - 1
}
