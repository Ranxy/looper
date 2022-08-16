package compilation

import (
	"sync/atomic"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/evaluator"
	"github.com/Ranxy/looper/syntax"
)

type Compilation struct {
	boundGlobalScope atomic.Pointer[bind.BoundGlobalScope]

	Previous   *Compilation
	SyntaxTree *syntax.SyntaxTree
}

func NewCompliation(previous *Compilation, syntaxTree *syntax.SyntaxTree) *Compilation {
	return &Compilation{
		Previous:   previous,
		SyntaxTree: syntaxTree,
	}
}
func (c *Compilation) ContinueWith(syntaxTree *syntax.SyntaxTree) *Compilation {
	return NewCompliation(c, syntaxTree)
}

func (c *Compilation) GlobalScope() *bind.BoundGlobalScope {
	if c == nil {
		return nil
	}
	gs := c.boundGlobalScope.Load()
	if gs == nil {
		gs = bind.BindGlobalScope(c.Previous.GlobalScope(), c.SyntaxTree.Root)
		c.boundGlobalScope.CompareAndSwap(nil, gs)
	}

	return gs
}

func (c *Compilation) Evaluate(variables map[syntax.VariableSymbol]any) *evaluator.EvaluateResult {
	c.GlobalScope().Diagnostic = c.GlobalScope().Diagnostic.Merge(c.SyntaxTree.Diagnostics)
	diagnostics := c.GlobalScope().Diagnostic
	if len(diagnostics.List) > 0 {
		return &evaluator.EvaluateResult{
			Diagnostic: diagnostics,
			Value:      nil,
		}
	}
	eval := evaluator.NewEvaluater(c.GlobalScope().Statements, variables)
	value := eval.Evaluate()
	return &evaluator.EvaluateResult{
		Diagnostic: diagnostics,
		Value:      value,
	}
}
