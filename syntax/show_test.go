package syntax

import (
	"testing"
)

func TestPrintExpress(t *testing.T) {

	expr := "(1 + 2) * (3 - 1)"
	p := NewParser(expr)
	tree := p.Parse()

	PrintExpress(tree.Root, "", true)
}
