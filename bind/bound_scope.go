package bind

import "github.com/Ranxy/looper/syntax"

type BoundScope struct {
	Parent    *BoundScope
	variables map[string]*syntax.VariableSymbol
}

func NewBoundScope(parent *BoundScope) *BoundScope {
	return &BoundScope{
		Parent:    parent,
		variables: map[string]*syntax.VariableSymbol{},
	}
}

func (s *BoundScope) TryDeclare(variable *syntax.VariableSymbol) bool {
	if _, has := s.variables[variable.Name]; has {
		return false
	}
	s.variables[variable.Name] = variable
	return true
}

func (s *BoundScope) TryLookup(name string) (*syntax.VariableSymbol, bool) {
	if v, has := s.variables[name]; has {
		return v, true
	}
	if s.Parent == nil {
		return nil, false
	}
	return s.Parent.TryLookup(name)
}

func (s *BoundScope) GetDeclareVariables() []*syntax.VariableSymbol {
	res := make([]*syntax.VariableSymbol, 0, len(s.variables))
	for _, v := range s.variables {
		res = append(res, v)
	}
	return res
}
