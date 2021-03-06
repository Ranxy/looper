package diagnostic

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Ranxy/looper/texts"
)

type Diagnostic struct {
	Span    texts.TextSpan
	Message string
}

func (d *Diagnostic) String() string {
	return d.Message
}

func (d *Diagnostic) StringWithLine(padding int, line string) string {
	if d.Span.End() > len(line) {
		return line
	}
	padding += d.Span.Start()

	sb := strings.Builder{}
	sb.WriteString(line)
	sb.WriteString("\033[31m")
	sb.WriteByte('\n')
	for i := 0; i < padding; i++ {
		sb.WriteByte(' ')
	}
	for i := d.Span.Start(); i < d.Span.End(); i++ {
		sb.WriteByte('^')
	}
	sb.WriteByte('\n')
	for i := 0; i < padding; i++ {
		sb.WriteByte(' ')
	}
	sb.WriteByte('|')
	sb.WriteByte('\n')
	for i := 0; i < padding; i++ {
		sb.WriteByte(' ')
	}
	sb.WriteString(d.Message)
	sb.WriteString("\033[0m")
	sb.WriteByte('\n')
	return sb.String()
}

type DiagnosticBag struct {
	List []Diagnostic
}

func NewDiagnostics() *DiagnosticBag {
	return &DiagnosticBag{
		List: make([]Diagnostic, 0),
	}
}
func MergeDiagnostics(b *DiagnosticBag) *DiagnosticBag {
	res := NewDiagnostics()
	res.Merge(b)
	return res
}

func (b *DiagnosticBag) Print(codeLine string) {
	for _, d := range b.List {
		fmt.Print(d.StringWithLine(0, codeLine))
	}
}
func (b *DiagnosticBag) PrintWithSource(source *texts.TextSource) {
	for _, d := range b.List {

		idx := source.GetLineIndex(d.Span.Start())
		line := source.Lines[idx]
		runeIdx := d.Span.Start() - line.Start + 1
		lineText := source.StringSpan(line.Span())
		lineOut := fmt.Sprintf("(%d, %d): ", idx+1, runeIdx)
		fmt.Printf(lineOut)
		fmt.Print(d.StringWithLine(len(lineOut), lineText))
	}
}

func (b *DiagnosticBag) Merge(bag *DiagnosticBag) {
	b.List = append(b.List, bag.List...)
}

func (b *DiagnosticBag) Report(span texts.TextSpan, message string) {
	b.List = append(b.List, Diagnostic{span, message})
}

func (b *DiagnosticBag) InvalidNumber(span texts.TextSpan, text string, tp reflect.Kind) {
	message := fmt.Sprintf("The number %s isn't valid %s.", text, tp)
	b.Report(span, message)
}
func (b *DiagnosticBag) BadCharacter(pos int, c rune) {
	span := texts.NewTextSpan(pos, 1)
	message := fmt.Sprintf("Bad character input: '%b'.", c)
	b.Report(span, message)
}

func (b *DiagnosticBag) UnexpectedToken(span texts.TextSpan, actualKind, expectedKind reflect.Kind) {
	message := fmt.Sprintf("Unexpected token %s, expected %s.", actualKind, expectedKind)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedUnaryOperator(span texts.TextSpan, operatorText string, operandType reflect.Kind) {
	message := fmt.Sprintf("Unary operator %s is not defined for type %s.", operatorText, operandType)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedBinaryOperator(span texts.TextSpan, operatorText string, leftType, rightType reflect.Kind) {
	message := fmt.Sprintf("Binary operator %s is not defined for type %s and %s.", operatorText, leftType, rightType)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedName(span texts.TextSpan, name string) {
	message := fmt.Sprintf("Variable '%s' doesn't exist.", name)
	b.Report(span, message)
}
