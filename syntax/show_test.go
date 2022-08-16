package syntax

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestPrintExpress(t *testing.T) {

	expr := "1 + 1"
	source := texts.NewTextSource([]rune(expr))
	tree := newSyntaxTree(source)

	err := PrintExpress(os.Stdout, tree.Root, "", true)
	require.NoError(t, err)
}
