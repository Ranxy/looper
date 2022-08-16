package syntax

import (
	"fmt"
	"io"

	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/texts"
)

type SyntaxTree struct {
	Text        *texts.TextSource
	Root        *CompliationUnit
	Diagnostics *diagnostic.DiagnosticBag
}

func newSyntaxTree(texts *texts.TextSource) *SyntaxTree {
	parser := NewParser(texts)
	root := parser.ParseComplitionUnit()
	return &SyntaxTree{
		Text:        texts,
		Root:        root,
		Diagnostics: diagnostic.MergeDiagnostics(parser.diagnostics),
	}
}

func ParseToTree(sourceText *texts.TextSource) *SyntaxTree {

	return newSyntaxTree(sourceText)
}

func ParseToTreeRaw(text string) *SyntaxTree {
	sourceText := texts.NewTextSource([]rune(text))
	return ParseToTree(sourceText)
}

func (s *SyntaxTree) Print(w io.Writer) (err error) {
	for _, err := range s.Diagnostics.List {
		_, e2 := w.Write([]byte(fmt.Sprintln("ERROR: ", err)))
		if e2 != nil {
			return e2
		}
	}

	err = PrintExpress(w, s.Root, "", true)
	if err != nil {
		return err
	}
	return nil
}
