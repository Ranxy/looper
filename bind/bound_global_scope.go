package bind

import (
	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/symbol"
)

type BoundGlobalScope struct {
	Previous   *BoundGlobalScope
	Diagnostic *diagnostic.DiagnosticBag
	Functions  []*symbol.FunctionSymbol
	Variables  []symbol.VariableSymbol
	Statements []Boundstatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope,
	diagnostic *diagnostic.DiagnosticBag,
	functions []*symbol.FunctionSymbol,
	variables []symbol.VariableSymbol,
	statements []Boundstatement) *BoundGlobalScope {

	return &BoundGlobalScope{
		Previous:   previous,
		Diagnostic: diagnostic,
		Functions:  functions,
		Variables:  variables,
		Statements: statements,
	}
}
