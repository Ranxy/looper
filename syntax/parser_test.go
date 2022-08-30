package syntax

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestCase(t *testing.T) {
	tests := []struct {
		text string
	}{
		{
			text: "1 + 2",
		},
		{
			text: "10 - 2 * 3",
		},
		{
			text: "1 + 2 * (3 -1 )",
		},
		{
			text: "-2",
		},
		{
			text: "1+-2",
		},
		{
			text: "-(1 * 2) -3",
		},
		{
			text: "1 < 2",
		},
		{
			text: "1 <= 2",
		},
		{
			text: "1 > 2",
		},
		{
			text: "1 >= 2",
		},
		{
			text: "1 & 2",
		},
		{
			text: "1 | 2",
		},
		{
			text: "1 ^ 2",
		},
		{
			text: "~1",
		},
		{
			text: "1 & 2 | 3 ^ ~ 4",
		},
		{
			text: `"hello" + "world"`,
		},
		{
			text: `{var a = "hello" a+ "world"}`,
		},
		{
			text: `{var a = ""}`,
		},
		{
			text: "{ var a = 0 if a == 2 a = 3 else a = 6 a }",
		},
		{
			text: "{ var a = 7 if a == 2 a = 3 a }",
		},
		{
			text: "{ var i = 10 var result = 0 while i != 0 { result = result + i i = i - 1} result }",
		},
		{
			text: "{ var result = 0 for var i = 1; i < 5; i=i+1 { result = result + i } result }",
		},
		{
			text: "{ var result = 0 for i = 1; i < 5; i=i+1 { result = result + i } result }",
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			source := texts.NewTextSource([]rune(tt.text))
			tree := newSyntaxTree(source)
			err := tree.Print(os.Stdout)
			if tree.Diagnostics.Has() {
				tree.Diagnostics.Print(tt.text)
			}
			require.NoError(t, err)
		})
	}
}

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

func TestNewParser_parser_if(t *testing.T) {
	expr := "if(a+1 == 2){var b = 0}else{var b = 3}"
	source := texts.NewTextSource([]rune(expr))
	tree := newSyntaxTree(source)
	require.Zero(t, len(tree.Diagnostics.List))
	err := tree.Print(os.Stdout)
	require.NoError(t, err)
}

func TestNewParser_parser_while(t *testing.T) {
	expr := "{ var i = 10 var result = 0 while i != 0 { result = result + i i = i - 1} result }"
	source := texts.NewTextSource([]rune(expr))
	tree := newSyntaxTree(source)
	require.Zero(t, len(tree.Diagnostics.List))
	err := tree.Print(os.Stdout)
	require.NoError(t, err)
}
