package syntax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	expr := "a=a+1 "
	p := NewParser(expr)
	if len(p.errors) != 0 {
		fmt.Println(p.errors)
		t.Fail()
	}
	fmt.Println(1)
	tree := p.Parse()
	if len(p.errors) != 0 {
		fmt.Println(p.errors)
		t.Fail()
	}
	tree.Print()
}
func TestNewParser_parser(t *testing.T) {
	expr := "1 + -2 * ---3"
	p := NewParser(expr)
	tree := p.Parse()
	require.Zero(t, len(tree.Errors))
	tree.Print()
}
func TestNewParser_parser_bool(t *testing.T) {
	expr := "(!true && !false ||  false ==false ) ==false && 1 == 1"
	p := NewParser(expr)
	tree := p.Parse()
	require.Zero(t, len(tree.Errors))
	tree.Print()
}
