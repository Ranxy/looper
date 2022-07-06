package syntax

import "fmt"

type SyntaxTree struct {
	Root   Express
	Eof    SyntaxToken
	Errors []string
}

func NewSyntaxTree(root Express, eof SyntaxToken, errors []string) *SyntaxTree {
	return &SyntaxTree{
		Root:   root,
		Eof:    eof,
		Errors: errors,
	}
}

func ParseToTree(text string) *SyntaxTree {
	p := NewParser(text)

	return p.Parse()
}

func (s *SyntaxTree) Eval() (int64, error) {
	return Evaluate(s.Root)
}

func (s *SyntaxTree) Print() {
	for _, err := range s.Errors {
		fmt.Println("ERROR: ", err)
	}

	PrintExpress(s.Root, "", true)
}
