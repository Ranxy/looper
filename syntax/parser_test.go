package syntax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	expr := "- 10 + 1 * 2   3 "
	p := NewParser(expr)
	fmt.Println(p.tokenList)
}
func TestNewParser_parser(t *testing.T) {
	expr := "1 + 2 * 3"
	p := NewParser(expr)
	tree := p.Parse()
	require.Zero(t, len(tree.Errors))
	tree.Print()
}
