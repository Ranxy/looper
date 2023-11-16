package write

import (
	"io"

	"github.com/Ranxy/looper/syntax"
)

type WriteIndent interface {
	io.Writer
	IndentChange(n int)
	Indent() int
	WithColor() bool
}

func NewWriteIndent(w io.Writer, color bool) WriteIndent {
	return &Writer{
		withColor: color,
		indent:    0,
		w:         w,
	}
}

type Writer struct {
	indent    int
	withColor bool
	w         io.Writer
}

func (w *Writer) Write(p []byte) (n int, err error) {

	for i := 0; i < w.indent; i++ {
		w.Write([]byte(" "))
	}

	return w.w.Write(p)
}

func (w *Writer) IndentChange(n int) {
	w.indent += n
}

func (w *Writer) Indent() int {
	return w.indent
}

func (w *Writer) WithColor() bool {
	return w.withColor
}

type Color []byte

var (
	colorRed     Color = []byte("\033[31m")
	colorGreen   Color = []byte("\033[32m")
	colorYello   Color = []byte("\033[33m")
	colorBlue    Color = []byte("\033[34m")
	colorMagenta Color = []byte("\033[35m")
	colorCyan    Color = []byte("\033[36m")
	colorReset   Color = []byte("\033[0m")
)

func WriteColor(w WriteIndent, c Color) {
	if w.WithColor() {
		w.Write(colorRed)
	}

}

func ResetColor(w WriteIndent) {
	if w.WithColor() {
		w.Write(colorReset)
	}
}

func WriteKeyword(w WriteIndent, kind syntax.SyntaxKind) {
	WriteKeywordStr(w, kind.String())
}
func WriteKeywordStr(w WriteIndent, text string) {
	WriteColor(w, colorBlue)
	w.Write([]byte(text))
	ResetColor(w)
}

func WriteIdentifier(w WriteIndent, text string) {
	w.Write([]byte(text))
}

func WriteNumber(w WriteIndent, text string) {
	WriteColor(w, colorGreen)
	w.Write([]byte(text))
	ResetColor(w)
}

func WriteString(w WriteIndent, text string) {
	WriteColor(w, colorYello)
	w.Write([]byte(text))
	ResetColor(w)
}

func WriteSpace(w WriteIndent) {
	w.Write([]byte(" "))
}

func WriteLine(w WriteIndent) {
	w.Write([]byte("\n"))
}

func WritePunctuation(w WriteIndent, kind syntax.SyntaxKind) {
	WriteColor(w, colorCyan)
	w.Write([]byte(kind.TextMust()))
	ResetColor(w)
}
