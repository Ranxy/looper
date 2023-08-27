package optimize

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestLowerRewrite_RewriteIfStatement(t *testing.T) {
	text := "{var b = 0{let a = 2+1 { if(-a == 1){ b = 2}else{ b = 3}}}}"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)

	boundTree := bind.BindGlobalScope(nil, tree.Root)
	if len(boundTree.Diagnostic.List) != 0 {
		boundTree.Diagnostic.Print(textSource)
		t.FailNow()
	}

	lowererStatement := Lower(bind.NewBoundBlockStatement(boundTree.Statements))

	err := bind.PrintBoundTree(os.Stdout, lowererStatement)
	require.NoError(t, err)
	require.Len(t, lowererStatement.Statement, 8) //2 declaration and a lowerer if statement
}

func TestLowerRewrite_BindWhileStatement(t *testing.T) {
	text := "{var result = 0 { var i = 10  while i != 0 { result = result + i i = i - 1}} result }"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)

	require.False(t, boundTree.Diagnostic.Has())

	lowerWhile := Lower(bind.NewBoundBlockStatement(boundTree.Statements))

	err := bind.PrintBoundTree(os.Stdout, lowerWhile)
	require.NoError(t, err)
}

func TestLowerRewrite_BindForStatement(t *testing.T) {
	text := `{
		var i = 0 
		var result = 0 
		for i = 1; i < 6 ;i = i + 1 {

			result = result + i
		} 
		result
	}`
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)

	require.False(t, boundTree.Diagnostic.Has())

	lowerWhile := Lower(bind.NewBoundBlockStatement(boundTree.Statements))

	err := bind.PrintBoundTree(os.Stdout, lowerWhile)
	require.NoError(t, err)
}
