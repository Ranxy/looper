package syntax

import (
	"fmt"
	"io"

	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/texts"
)

type SyntaxTree struct {
	Text        *texts.TextSource
	Root        Express
	Eof         SyntaxToken
	Diagnostics *diagnostic.DiagnosticBag
}

func NewSyntaxTree(root Express, eof SyntaxToken, source *texts.TextSource, diagnostics *diagnostic.DiagnosticBag) *SyntaxTree {
	return &SyntaxTree{
		Text:        source,
		Root:        root,
		Eof:         eof,
		Diagnostics: diagnostic.MergeDiagnostics(diagnostics),
	}
}

func ParseToTree(sourceText *texts.TextSource) *SyntaxTree {
	p := NewParser(sourceText)
	return p.Parse()
}

func ParseToTreeRaw(text string) *SyntaxTree {
	sourceText := texts.NewTextSource([]rune(text))
	p := NewParser(sourceText)
	return p.Parse()
}

func (s *SyntaxTree) Print(w io.Writer) {
	for _, err := range s.Diagnostics.List {
		w.Write([]byte(fmt.Sprintln("ERROR: ", err)))
	}

	PrintExpress(w, s.Root, "", true)
}
