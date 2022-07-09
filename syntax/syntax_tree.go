package syntax

import (
	"fmt"

	"github.com/Ranxy/looper/diagnostic"
)

type SyntaxTree struct {
	Root        Express
	Eof         SyntaxToken
	Diagnostics *diagnostic.DiagnosticBag
}

func NewSyntaxTree(root Express, eof SyntaxToken, diagnostics *diagnostic.DiagnosticBag) *SyntaxTree {
	return &SyntaxTree{
		Root:        root,
		Eof:         eof,
		Diagnostics: diagnostic.MergeDiagnostics(diagnostics),
	}
}

func ParseToTree(text string) *SyntaxTree {
	p := NewParser(text)

	return p.Parse()
}

func (s *SyntaxTree) Print() {
	for _, err := range s.Diagnostics.List {
		fmt.Println("ERROR: ", err)
	}

	PrintExpress(s.Root, "", true)
}
