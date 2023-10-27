package bind_test

import (
	"os"
	"strings"
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/bind/program"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestCreateControlFlow(t *testing.T) {
	text := `{
		var result = 0
		for var i = 0; i < 100; i=i+1 {
			result = result + i
		}
		var z = 0
		result
	}`
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)

	require.False(t, tree.Diagnostics.Has())
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	require.False(t, boundTree.Diagnostic.Has())
	program := program.BindProgram(boundTree)

	gf := bind.CreateControlFlow(program.Statement)
	sb := strings.Builder{}

	gf.WriteTo(&sb)

	err := os.WriteFile("cc.dot", []byte(sb.String()), os.ModePerm)
	require.NoError(t, err)
}
