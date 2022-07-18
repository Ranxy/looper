package syntax

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	expr := "a=2+1 - c + 2 * 3 "
	source := texts.NewTextSource([]rune(expr))
	p := NewParser(source)
	p.diagnostics.Print(expr)
	tree := p.Parse()
	p.diagnostics.Print(expr)
	tree.Print(os.Stdout)
}
func TestNewParser_parser(t *testing.T) {
	expr := "1 + -2 * ---3"
	source := texts.NewTextSource([]rune(expr))
	p := NewParser(source)
	tree := p.Parse()
	require.Zero(t, len(tree.Diagnostics.List))
	tree.Print(os.Stdout)
}
func TestNewParser_parser_bool(t *testing.T) {
	expr := "(!true && !false ||  false ==false ) ==false && 1 == 1"
	source := texts.NewTextSource([]rune(expr))
	p := NewParser(source)
	tree := p.Parse()
	require.Zero(t, len(tree.Diagnostics.List))
	tree.Print(os.Stdout)
}
