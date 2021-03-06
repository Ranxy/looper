package bind

import (
	"reflect"
	"testing"

	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinder_BindExpression(t *testing.T) {
	vm := NewVariableManage()

	text := "a = 2"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.NewParser(textSource).Parse()
	bound := NewBinder(vm)
	boundTree := bound.BindExpression(tree.Root)
	bound.Diagnostics.Print(text)
	require.Zero(t, len(bound.Diagnostics.List))

	assert.Equal(t, BoundNodeKindAssignmentExpress, boundTree.Kind())
	symbolA := vm.GetSymbol("a")
	assert.NotNil(t, symbolA)
	assert.Equal(t, reflect.Int64, symbolA.Type)
}
