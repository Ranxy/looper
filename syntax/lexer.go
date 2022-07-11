package syntax

import (
	"reflect"
	"strconv"
	"unicode"

	"github.com/Ranxy/looper/diagnostic"
)

type Lexer struct {
	text []rune
	len  int
	pos  int

	_start int
	_kind  SyntaxKind
	_value any

	diagnostics *diagnostic.DiagnosticBag
}

func NewLexer(text string) *Lexer {
	list := []rune(text)
	return &Lexer{
		text:        list,
		len:         len(list),
		pos:         0,
		diagnostics: diagnostic.NewDiagnostics(),
	}
}

func (l *Lexer) Diagnostics() *diagnostic.DiagnosticBag {
	return l.diagnostics
}

func (l *Lexer) Current() rune {
	return l.Peek(0)
}
func (l *Lexer) Lookahead() rune {
	return l.Peek(1)
}

func (l *Lexer) Peek(offset int) rune {
	idx := l.pos + offset

	if idx >= l.len {
		return unicode.MaxRune
	}

	return l.text[idx]
}

func (l *Lexer) next(n int) {
	l.pos += n
}

func (l *Lexer) NextToken() SyntaxToken {
	l._start = l.pos
	l._kind = SyntaxKindBadToken
	l._value = nil

	switch l.Current() {
	case unicode.MaxRune:
		l._kind = SyntaxKindEofToken
	case '+':
		l._kind = SyntaxKindPlusToken
		l.next(1)
	case '-':
		l._kind = SyntaxKindMinusToken
		l.next(1)
	case '*':
		l._kind = SyntaxKindStarToken
		l.next(1)
	case '/':
		l._kind = SyntaxKindSlashToken
		l.next(1)
	case '(':
		l._kind = SyntaxKindOpenParenthesisToken
		l.next(1)
	case ')':
		l._kind = SyntaxKindCloseParenthesisToken
		l.next(1)

	case '&':
		if l.Lookahead() == '&' {
			l._kind = SyntaxKindAmpersandAmpersandToken
			l.next(2)
		}
	case '|':
		if l.Lookahead() == '|' {
			l._kind = SyntaxKindPipePileToken
			l.next(2)
		}
	case '=':
		if l.Lookahead() == '=' {
			l._kind = SyntaxKindEqualEqualToken
			l.next(2)
		} else {
			l._kind = SyntaxKindEqualToken
			l.next(1)
		}
	case '!':
		if l.Lookahead() == '=' {
			l._kind = SyntaxKindBangEqualToken
			l.next(2)
		} else {
			l._kind = SyntaxKindBangToken
			l.next(1)
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.readDigit()
	case ' ', '\t', '\n', '\r':
		l.readSpace()
	default:
		if unicode.IsLetter(l.Current()) {
			l.readIdentifier()
		} else if unicode.IsSpace(l.Current()) {
			l.readSpace()
		} else {
			l.diagnostics.BadCharacter(l.pos, l.Current())
			l.next(1)
		}
	}

	text := l._kind.Text()
	if text == "" {
		text = string(l.text[l._start:l.pos])
	}

	return SyntaxToken{l._kind, l._start, text, l._value}
}

func (l *Lexer) readDigit() {
	for unicode.IsDigit(l.Current()) {
		l.next(1)
	}

	text := string(l.text[l._start:l.pos])
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		l.diagnostics.InvalidNumber(diagnostic.NewTextSpan(l._start, l.pos-l._start), text, reflect.Int64)
	}

	l._value = v
	l._kind = SyntaxKindNumberToken
}

func (l *Lexer) readSpace() {
	for unicode.IsSpace(l.Current()) {
		l.next(1)
	}
	l._kind = SyntaxKindWhiteSpaceToken
}

func (l *Lexer) readIdentifier() {
	for unicode.IsLetter(l.Current()) {
		l.next(1)
	}
	text := l.text[l._start:l.pos]
	l._kind = GetKeyWordsKind(string(text))
}

func ParseTokens(text string) (res []SyntaxToken) {

	lex := NewLexer(text)
	for {
		token := lex.NextToken()
		if token.Kind() == SyntaxKindEofToken {
			break
		}
		res = append(res, token)
	}
	return res
}
