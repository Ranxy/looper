package bind

import (
	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/syntax"
)

type BoundGlobalScope struct {
	Previous   *BoundGlobalScope
	Diagnostic *diagnostic.DiagnosticBag
	Variables  []*syntax.VariableSymbol
	Statements Boundstatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope,
	diagnostic *diagnostic.DiagnosticBag, variables []*syntax.VariableSymbol, statement Boundstatement) *BoundGlobalScope {

	return &BoundGlobalScope{
		Previous:   previous,
		Diagnostic: diagnostic,
		Variables:  variables,
		Statements: statement,
	}
}
