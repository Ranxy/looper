package bind

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Ranxy/looper/syntax"
)

type variable struct {
	Type  reflect.Kind
	Value any
}

type VariableManage map[string]*variable

func NewVariableManage() VariableManage {
	return make(VariableManage)
}
func (m VariableManage) GetSymbol(text string) *syntax.VariableSymbol {
	v, has := m[text]
	if !has {
		return nil
	}
	return &syntax.VariableSymbol{
		Name: text,
		Type: v.Type,
	}
}
func (m VariableManage) GetValue(text string) *variable {
	v, has := m[text]
	if !has {
		return nil
	}
	return v
}

func (m VariableManage) Add(variableSymbol *syntax.VariableSymbol, value any) {
	m[variableSymbol.Name] = &variable{
		Type:  variableSymbol.Type,
		Value: value,
	}
}
func (m VariableManage) Declare(variableSymbol *syntax.VariableSymbol) {
	if _, has := m[variableSymbol.Name]; !has {
		m[variableSymbol.Name] = &variable{
			Type:  variableSymbol.Type,
			Value: nil,
		}
	}
}

func (m VariableManage) Dump() string {
	sb := strings.Builder{}
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString(":")
		sb.WriteString(v.Type.String())
		sb.WriteString("->")
		sb.WriteString(fmt.Sprintf("%v", v.Value))
		sb.WriteString("\n")
	}
	return sb.String()
}
