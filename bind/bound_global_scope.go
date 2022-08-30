package bind

import (
	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/symbol"
)

type BoundGlobalScope struct {
	Previous   *BoundGlobalScope
	Diagnostic *diagnostic.DiagnosticBag
	Variables  []*symbol.VariableSymbol
	Statements Boundstatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope,
	diagnostic *diagnostic.DiagnosticBag, variables []*symbol.VariableSymbol, statement Boundstatement) *BoundGlobalScope {

	return &BoundGlobalScope{
		Previous:   previous,
		Diagnostic: diagnostic,
		Variables:  variables,
		Statements: statement,
	}
}
