package symbol

import (
	"strings"
)

type FunctionSymbol struct {
	name        string
	Parameter   []*ParameterSymbol
	Type        *TypeSymbol
	Declaration any
}

func NewFunctionSymbol(name string, parameter []*ParameterSymbol, tp *TypeSymbol, declaration any) *FunctionSymbol {
	return &FunctionSymbol{
		name:        name,
		Parameter:   parameter,
		Type:        tp,
		Declaration: declaration,
	}
}

func (s *FunctionSymbol) Kind() SymbolKind {
	return SymbolKindFunction
}

func (s *FunctionSymbol) GetName() string {
	return s.name
}

func (s *FunctionSymbol) String() string {
	sb := strings.Builder{}

	sb.WriteString(s.name)
	sb.WriteByte('(')
	paramLen := len(s.Parameter)
	for i, p := range s.Parameter {
		sb.WriteString(p.String())
		if i != paramLen-1 {
			sb.WriteByte(',')
		}
	}
	sb.WriteString(")->")
	sb.WriteString(s.Type.GetName())
	return sb.String()
}
