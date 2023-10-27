package diagnostic

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/texts"
)

type Diagnostic struct {
	Span    texts.TextSpan
	Message string
}

func (d *Diagnostic) String() string {
	return d.Message
}

func (d *Diagnostic) StringWithLine(padding int, line texts.TextLine) string {
	if d.Span.End() > line.End() {
		return line.String()
	}
	lineStart := d.Span.Start() - line.Start
	tabCount := line.TabCount(lineStart)
	padding = padding + lineStart - tabCount

	sb := strings.Builder{}
	sb.WriteString(line.String())
	sb.WriteString("\033[31m")
	sb.WriteByte('\n')

	printPadding := func() {
		for i := 0; i < tabCount; i++ {
			sb.WriteByte('\t')
		}
		for i := 0; i < padding; i++ {
			sb.WriteByte(' ')
		}
	}

	printPadding()
	for i := d.Span.Start(); i < d.Span.End(); i++ {
		sb.WriteByte('^')
	}
	sb.WriteByte('\n')
	printPadding()
	sb.WriteByte('|')
	sb.WriteByte('\n')
	printPadding()
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
	res = res.Merge(b)
	return res
}
func (b *DiagnosticBag) Has() bool {
	return len(b.List) != 0
}

func (b *DiagnosticBag) Print(source *texts.TextSource) {
	b.PrintWithSourceStdout(source)
}

func (b *DiagnosticBag) PrintWithSourceStdout(source *texts.TextSource) {
	b.PrintWithSource(os.Stdout, source)
}

func (b *DiagnosticBag) PrintWithSource(w io.Writer, source *texts.TextSource) {
	b.Sort()

	for _, d := range b.List {

		idx := source.GetLineIndex(d.Span.Start())
		line := source.Lines[idx]
		runeIdx := d.Span.Start() - line.Start + 1
		lineOut := fmt.Sprintf("(%d, %d): ", idx+1, runeIdx)
		w.Write([]byte(lineOut))
		w.Write([]byte(d.StringWithLine(len(lineOut), line)))
	}
}

func (b *DiagnosticBag) Sort() {
	sort.Slice(b.List, func(i, j int) bool {
		cmp := b.List[i].Span.Start() - b.List[j].Span.Start()
		if cmp == 0 {
			cmp = b.List[i].Span.Length() - b.List[j].Span.Length()
		}

		return cmp < 0
	})
}

func (b *DiagnosticBag) Reset() {
	b.List = b.List[0:0]
}

func (b *DiagnosticBag) Merge(bag *DiagnosticBag) *DiagnosticBag {
	list := append(b.List, bag.List...)
	return &DiagnosticBag{List: list}
}

func (b *DiagnosticBag) Report(span texts.TextSpan, message string) {
	b.List = append(b.List, Diagnostic{span, message})
}

func (b *DiagnosticBag) InvalidNumber(span texts.TextSpan, text string, tp *symbol.TypeSymbol) {
	message := fmt.Sprintf("The number %s isn't valid %s.", text, tp)
	b.Report(span, message)
}
func (b *DiagnosticBag) BadCharacter(pos int, c rune) {
	span := texts.NewTextSpan(pos, 1)
	message := fmt.Sprintf("Bad character input: '%b'.", c)
	b.Report(span, message)
}

func (b *DiagnosticBag) UnexpectedToken(span texts.TextSpan, actualKind, expectedKind *symbol.TypeSymbol) {
	message := fmt.Sprintf("Unexpected token %s, expected %s.", actualKind, expectedKind)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedUnaryOperator(span texts.TextSpan, operatorText string, operandType *symbol.TypeSymbol) {
	message := fmt.Sprintf("Unary operator %s is not defined for type %s.", operatorText, operandType)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedBinaryOperator(span texts.TextSpan, operatorText string, leftType, rightType *symbol.TypeSymbol) {
	message := fmt.Sprintf("Binary operator %s is not defined for type %s and %s.", operatorText, leftType, rightType)
	b.Report(span, message)
}

func (b *DiagnosticBag) UndefinedName(span texts.TextSpan, name string) {
	message := fmt.Sprintf("Variable '%s' doesn't exist.", name)
	b.Report(span, message)
}

func (b *DiagnosticBag) CannotAssign(span texts.TextSpan, name string) {
	message := fmt.Sprintf("Variable '%s' is readonly and can not assign.", name)
	b.Report(span, message)
}

func (b *DiagnosticBag) CannotConvert(span texts.TextSpan, expect, actual *symbol.TypeSymbol) {
	message := fmt.Sprintf("Cannot convert '%s' to '%s'.", expect, actual)
	b.Report(span, message)
}

func (b *DiagnosticBag) VariableAlreadyDeclared(span texts.TextSpan, name string) {
	message := fmt.Sprintf("Variable '%s' is already declared.", name)
	b.Report(span, message)
}

func (b *DiagnosticBag) ReportUnterminatedString(span texts.TextSpan) {
	b.Report(span, "Unterminated string literal.")
}

func (b *DiagnosticBag) ExpressionMustReturnValue(span texts.TextSpan) {
	b.Report(span, "Expression must return a value.")
}

func (b *DiagnosticBag) UndefinedFunction(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("Function %s not defined", name)
	b.Report(span, msg)
}

func (b *DiagnosticBag) FunctionNotTopLevel(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("Function %s not at top level", name)
	b.Report(span, msg)
}

func (b *DiagnosticBag) UndefinedType(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("Type %s not defined", name)
	b.Report(span, msg)
}

func (b *DiagnosticBag) WrongArgumentNumber(span texts.TextSpan, name string, expect, actual int) {
	msg := fmt.Sprintf("Function %s requires %d arguments, but was given %d", name, expect, actual)
	b.Report(span, msg)
}

func (b *DiagnosticBag) WrongArgumentType(span texts.TextSpan, name string, expect, actual *symbol.TypeSymbol) {
	msg := fmt.Sprintf("Paramenter %s require a value of %s, but was given a value of %s",
		name, expect.String(), actual.String())
	b.Report(span, msg)
}
func (b *DiagnosticBag) ParameterAlreadyDeclared(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("Paramenter %s already declared", name)
	b.Report(span, msg)
}

func (b *DiagnosticBag) SymbolAlreadyDeclared(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("Symbol %s already declared", name)
	b.Report(span, msg)
}

func (b *DiagnosticBag) ReportInvalidBreakOrContinue(span texts.TextSpan, name string) {
	msg := fmt.Sprintf("The keyword %s can only be used inside of loops", name)
	b.Report(span, msg)
}
func (b *DiagnosticBag) ReportInvalidReturn(span texts.TextSpan) {
	msg := "The keyword return can only be used inside of function"
	b.Report(span, msg)
}

func (b *DiagnosticBag) ReportAllPathMustReturn(span texts.TextSpan) {
	msg := "Not all code paths return a value"
	b.Report(span, msg)
}
