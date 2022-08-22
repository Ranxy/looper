package optimize

import (
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestFlattenStatement(t *testing.T) {
	text := "{var result = 0 {var i = 0 {for i = 1; i < 5 ;i = i + 1 { result = result + 1 }} result }}"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)

	boundTree := bind.BindGlobalScope(nil, tree.Root)
	if len(boundTree.Diagnostic.List) != 0 {
		boundTree.Diagnostic.Print(text)
		t.FailNow()
	}
	befor := boundTree.Statements.(*bind.BoundBlockStatements)
	require.Equal(t, 2, len(befor.Statement))
	require.Len(t, boundTree.Diagnostic.List, 0)
	BlockStatement := FlattenStatement(boundTree.Statements)

	require.Equal(t, 4, len(BlockStatement.Statement))
}
