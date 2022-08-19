package bind

import (
	"testing"

	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestBinder_BindExpression(t *testing.T) {
	text := "let a = 2"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := BindGlobalScope(nil, tree.Root)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)
	require.Len(t, boundTree.Variables, 1)
	require.Equal(t, boundTree.Variables[0].Name, "a")
	require.Equal(t, boundTree.Variables[0].IsReadOnly, true)
}

func TestBinder_BindifStatement(t *testing.T) {
	text := "let a = 2; if(-a == 1){ b = 2}else{ b = 3}"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := BindGlobalScope(nil, tree.Root)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)

	require.False(t, boundTree.Diagnostic.Has())
}
