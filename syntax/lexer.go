package syntax

import (
	"fmt"
	"strconv"
	"unicode"
)

type Lexer struct {
	text []rune
	len  int
	pos  int

	errors []string
}

func NewLexer(text string) *Lexer {
	list := []rune(text)
	return &Lexer{
		text:   list,
		len:    len(list),
		pos:    0,
		errors: []string{},
	}
}

func (l *Lexer) Error() []string {
	return l.errors
}

func (l *Lexer) Current() rune {
	if l.pos >= l.len {
		return unicode.MaxRune
	}

	return l.text[l.pos]
}

func (l *Lexer) posAndNext() int {
	pos := l.pos
	l.pos += 1
	return pos
}

func (l *Lexer) next() {
	l.pos += 1
}

func (l *Lexer) NextToken() SyntaxToken {
	if l.pos >= l.len {
		return SyntaxToken{
			kind:     SyntaxKindEofToken,
			Position: l.pos,
			Text:     "",
			Value:    nil,
		}
	}

	if unicode.IsDigit(l.Current()) {
		start := l.pos
		for unicode.IsDigit(l.Current()) {
			l.next()
		}

		text := string(l.text[start:l.pos])
		v, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			l.errors = append(l.errors, fmt.Sprintf("Number %s not a valid int", text))
		}

		return SyntaxToken{
			kind:     SyntaxKindNumberToken,
			Position: start,
			Text:     text,
			Value:    v,
		}
	}

	if unicode.IsSpace(l.Current()) {
		start := l.pos

		for unicode.IsSpace(l.Current()) {
			l.next()
		}

		text := string(l.text[start:l.pos])

		return SyntaxToken{
			kind:     SyntaxKindWhiteSpaceToken,
			Position: start,
			Text:     text,
			Value:    nil,
		}
	}

	//match predefine opt

	switch l.Current() {
	case '+':
		return SyntaxToken{SyntaxKindPlusToken, l.posAndNext(), "+", nil}
	case '-':
		return SyntaxToken{SyntaxKindMinusToken, l.posAndNext(), "-", nil}
	case '*':
		return SyntaxToken{SyntaxKindStarToken, l.posAndNext(), "*", nil}
	case '/':
		return SyntaxToken{SyntaxKindSlashToken, l.posAndNext(), "/", nil}
	case '(':
		return SyntaxToken{SyntaxKindOpenParenthesisToken, l.posAndNext(), "(", nil}
	case ')':
		return SyntaxToken{SyntaxKindCloseParenthesisToken, l.posAndNext(), ")", nil}
	}

	l.errors = append(l.errors, fmt.Sprintf("Bad Token input %v", l.Current()))

	return SyntaxToken{SyntaxKindBadToken, l.posAndNext(), string(l.text[l.pos:l.pos]), nil}
}
