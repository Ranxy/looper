package syntax

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/texts"
)

func TestPrintExpress(t *testing.T) {

	expr := "1 + 1"
	source := texts.NewTextSource([]rune(expr))
	p := NewParser(source)
	tree := p.Parse()

	PrintExpress(os.Stdout, tree.Root, "", true)
}
