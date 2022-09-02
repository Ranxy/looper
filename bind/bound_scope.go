package bind

import (
	"github.com/Ranxy/looper/symbol"
)

type BoundScope struct {
	Parent    *BoundScope
	variables map[string]*symbol.VariableSymbol
	functions map[string]*symbol.FunctionSymbol
}

func NewBoundScope(parent *BoundScope) *BoundScope {
	return &BoundScope{
		Parent:    parent,
		variables: map[string]*symbol.VariableSymbol{},
		functions: map[string]*symbol.FunctionSymbol{},
	}
}

func (s *BoundScope) TryDeclareVariable(variable *symbol.VariableSymbol) bool {
	if _, has := s.variables[variable.Name]; has {
		return false
	}
	s.variables[variable.Name] = variable
	return true
}

func (s *BoundScope) TryLookupVariable(name string) (*symbol.VariableSymbol, bool) {
	if v, has := s.variables[name]; has {
		return v, true
	}
	if s.Parent == nil {
		return nil, false
	}
	return s.Parent.TryLookupVariable(name)
}

func (s *BoundScope) GetDeclareVariables() []*symbol.VariableSymbol {
	res := make([]*symbol.VariableSymbol, 0, len(s.variables))
	for _, v := range s.variables {
		res = append(res, v)
	}
	return res
}

func (s *BoundScope) TryDeclareFunction(function *symbol.FunctionSymbol) bool {
	if _, has := s.functions[function.GetName()]; has {
		return false
	}
	s.functions[function.GetName()] = function
	return true
}

func (s *BoundScope) TryLookupFunction(name string) (*symbol.FunctionSymbol, bool) {
	if v, has := s.functions[name]; has {
		return v, true
	}
	if s.Parent == nil {
		return nil, false
	}
	return s.Parent.TryLookupFunction(name)
}

func (s *BoundScope) GetDeclareFunctions() []*symbol.FunctionSymbol {
	res := make([]*symbol.FunctionSymbol, 0, len(s.functions))
	for _, v := range s.functions {
		res = append(res, v)
	}
	return res
}
