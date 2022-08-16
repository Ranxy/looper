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
	tree := newSyntaxTree(source)
	err := tree.Print(os.Stdout)
	require.NoError(t, err)
}
func TestNewParser_parser(t *testing.T) {
	expr := "1 + -2 * ---3"
	source := texts.NewTextSource([]rune(expr))
	tree := newSyntaxTree(source)
	require.Zero(t, len(tree.Diagnostics.List))
	err := tree.Print(os.Stdout)
	require.NoError(t, err)
}
func TestNewParser_parser_bool(t *testing.T) {
	expr := "(!true && !false ||  false ==false ) ==false && 1 == 1"
	source := texts.NewTextSource([]rune(expr))
	tree := newSyntaxTree(source)
	require.Zero(t, len(tree.Diagnostics.List))
	err := tree.Print(os.Stdout)
	require.NoError(t, err)
}
