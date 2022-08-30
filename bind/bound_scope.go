package bind

import (
	"github.com/Ranxy/looper/symbol"
)

type BoundScope struct {
	Parent    *BoundScope
	variables map[string]*symbol.VariableSymbol
}

func NewBoundScope(parent *BoundScope) *BoundScope {
	return &BoundScope{
		Parent:    parent,
		variables: map[string]*symbol.VariableSymbol{},
	}
}

func (s *BoundScope) TryDeclare(variable *symbol.VariableSymbol) bool {
	if _, has := s.variables[variable.Name]; has {
		return false
	}
	s.variables[variable.Name] = variable
	return true
}

func (s *BoundScope) TryLookup(name string) (*symbol.VariableSymbol, bool) {
	if v, has := s.variables[name]; has {
		return v, true
	}
	if s.Parent == nil {
		return nil, false
	}
	return s.Parent.TryLookup(name)
}

func (s *BoundScope) GetDeclareVariables() []*symbol.VariableSymbol {
	res := make([]*symbol.VariableSymbol, 0, len(s.variables))
	for _, v := range s.variables {
		res = append(res, v)
	}
	return res
}
