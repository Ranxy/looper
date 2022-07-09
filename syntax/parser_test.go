package syntax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	expr := "a=2+1 - c + 2 * 3 "
	p := NewParser(expr)
	p.diagnostics.Print(expr)
	tree := p.Parse()
	p.diagnostics.Print(expr)
	tree.Print()
}
func TestNewParser_parser(t *testing.T) {
	expr := "1 + -2 * ---3"
	p := NewParser(expr)
	tree := p.Parse()
	require.Zero(t, len(tree.Diagnostics.List))
	tree.Print()
}
func TestNewParser_parser_bool(t *testing.T) {
	expr := "(!true && !false ||  false ==false ) ==false && 1 == 1"
	p := NewParser(expr)
	tree := p.Parse()
	require.Zero(t, len(tree.Diagnostics.List))
	tree.Print()
}
