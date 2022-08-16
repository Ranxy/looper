package evaluator

import "github.com/Ranxy/looper/diagnostic"

type EvaluateResult struct {
	Diagnostic *diagnostic.DiagnosticBag
	Value      any
}
