package program

import (
	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/optimize"
	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

func BindProgram(globalScope *bind.BoundGlobalScope) *BoundProgram {
	parentScope := bind.CreateParentScope(globalScope)
	functionBodys := make(map[*symbol.FunctionSymbol]*bind.BoundBlockStatements)
	diag := globalScope.Diagnostic

	scope := globalScope
	for scope != nil {
		for _, fn := range scope.Functions {
			binder := bind.NewBinder(parentScope, fn)
			body := binder.BindStatement(fn.Declaration.(*syntax.FunctionDeclarationSyntax).Body)
			lowerBody := optimize.Lower(body)
			functionBodys[fn] = lowerBody

			diag = diag.Merge(binder.Diagnostics)
		}

		scope = scope.Previous
	}

	statement := optimize.Lower(bind.NewBoundBlockStatement(globalScope.Statements))

	return &BoundProgram{
		Diagnostic: diag,
		Functions:  functionBodys,
		Statement:  statement,
	}
}

type BoundProgram struct {
	Diagnostic *diagnostic.DiagnosticBag
	Functions  map[*symbol.FunctionSymbol]*bind.BoundBlockStatements
	Statement  *bind.BoundBlockStatements
}
