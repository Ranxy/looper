package bind

import (
	"github.com/Ranxy/looper/symbol"
)

type BoundScope struct {
	Parent  *BoundScope
	symbols map[string]symbol.Symbol
}

func NewBoundScope(parent *BoundScope) *BoundScope {
	return &BoundScope{
		Parent:  parent,
		symbols: make(map[string]symbol.Symbol),
	}
}

func tryDeclare[T symbol.Symbol](s *BoundScope, tSymbol T) bool {
	if _, has := s.symbols[tSymbol.GetName()]; has {
		return false
	}
	s.symbols[tSymbol.GetName()] = tSymbol
	return true
}

func tryLookup[T symbol.Symbol](s *BoundScope, name string) (T, bool) {
	if v, has := s.symbols[name]; has {
		if res, ok := v.(T); ok {
			return res, true
		}
	}
	if s.Parent == nil {
		var zero T
		return zero, false
	}
	return tryLookup[T](s.Parent, name)
}

// func getDeclare[T symbol.Symbol](s *BoundScope) []T {
// 	res := make([]T, 0)
// 	for _, v := range s.symbols {
// 		if t, ok := v.(T); ok {
// 			res = append(res, t) //at go1.19,this will got ice https://github.com/golang/go/issues/54302 , will be update if golang fix this.
// 		}
// 	}
// 	return res
// }

func (s *BoundScope) TryDeclareVariable(variable symbol.VariableSymbol) bool {
	return tryDeclare(s, variable)
}

func (s *BoundScope) TryLookupVariable(name string) (symbol.VariableSymbol, bool) {
	return tryLookup[symbol.VariableSymbol](s, name)
}

func (s *BoundScope) GetDeclareVariables() []symbol.VariableSymbol {
	res := make([]symbol.VariableSymbol, 0)
	for _, v := range s.symbols {
		if t, ok := v.(symbol.VariableSymbol); ok {
			res = append(res, t)
		}
	}
	return res
}

func (s *BoundScope) TryDeclareFunction(function *symbol.FunctionSymbol) bool {
	return tryDeclare(s, function)
}

func (s *BoundScope) TryLookupFunction(name string) (*symbol.FunctionSymbol, bool) {
	return tryLookup[*symbol.FunctionSymbol](s, name)
}

func (s *BoundScope) GetDeclareFunctions() []*symbol.FunctionSymbol {
	res := make([]*symbol.FunctionSymbol, 0)
	for _, v := range s.symbols {
		if t, ok := v.(*symbol.FunctionSymbol); ok {
			res = append(res, t)
		}
	}
	return res
}
