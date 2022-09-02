package bind

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type BoundNodeKind int

const (
	BoundNodeKindErrorExpress BoundNodeKind = iota
	BoundNodeKindLiteralExpress
	BoundNodeKindVariableExpress
	BoundNodeKindAssignmentExpress
	BoundNodeKindUnaryExpress
	BoundNodeKindBinaryExpress
	BoundNodeKindCallExpress

	BoundNodeKindBlockStatement
	BoundNodeKindVariableDeclaration
	BoundNodeKindIfStatement
	BoundNodeKindWhileStatement
	BoundNodeKindForStatement
	BoundNodeKindLabelStatement
	BoundNodeKindGotoStatement
	BoundNodeKindConditionalGotoStatement
	BoundNodeKindExpressionStatement
)

var boundNodeKindNameMap = map[BoundNodeKind]string{
	BoundNodeKindLiteralExpress:    "LiteralExpress",
	BoundNodeKindVariableExpress:   "VariableExpress",
	BoundNodeKindAssignmentExpress: "AssignmentExpress",
	BoundNodeKindUnaryExpress:      "UnaryExpress",
	BoundNodeKindBinaryExpress:     "BinaryExpress",
	BoundNodeKindCallExpress:       "CallExpress",

	BoundNodeKindBlockStatement:           "BlockStatement",
	BoundNodeKindVariableDeclaration:      "VariableDeclaration",
	BoundNodeKindIfStatement:              "IfStatement",
	BoundNodeKindWhileStatement:           "WhileStatement",
	BoundNodeKindForStatement:             "ForStatement",
	BoundNodeKindLabelStatement:           "LabelStatement",
	BoundNodeKindGotoStatement:            "GotoStatement",
	BoundNodeKindConditionalGotoStatement: "ConditionalGotoStatement",
	BoundNodeKindExpressionStatement:      "ExpressionStatement",
}

func (b BoundNodeKind) String() string {
	str, has := boundNodeKindNameMap[b]
	if has {
		return str
	} else {
		return fmt.Sprintf("UnexceptedKind %d", b)
	}
}

type boolStringer bool

func (b *boolStringer) String() string {
	return strconv.FormatBool(bool(*b))
}
func newBookStringer(b bool) *boolStringer {
	x := boolStringer(b)
	return &x
}

type literalValue struct {
	v any
}

func (s *literalValue) String() string {
	return fmt.Sprintf("%v", s.v)
}

func PrintBoundTree(w io.Writer, node BoundNode) error {
	return prettyPrint(w, node, "", true)
}

func prettyPrint(w io.Writer, node BoundNode, indent string, isLast bool) error {
	var mark string
	if isLast {
		mark = "└──"
	} else {
		mark = "├──"
	}

	_, err := w.Write([]byte(indent))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(mark))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(NodeText(node)))
	if err != nil {
		return err
	}
	if node == nil {
		return nil
	}
	properties := node.GetProperties()
	if len(properties) > 0 {
		_, err = w.Write([]byte("{ "))
		if err != nil {
			return err
		}
		for i, p := range node.GetProperties() {
			if i != 0 {
				_, err = w.Write([]byte{','})
				if err != nil {
					return err
				}
			}
			_, err = w.Write([]byte(" "))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte(reflect.TypeOf(p).Elem().Name()))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("="))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte(p.String()))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte(" }"))
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	if isLast {
		indent += "   "
	} else {
		indent += "│  "
	}

	childrenList := node.GetChildren()
	if len(childrenList) != 0 {
		last := childrenList[len(childrenList)-1]
		for _, childen := range childrenList {
			err = prettyPrint(w, childen, indent, last == childen)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NodeText(node BoundNode) string {
	switch v := node.(type) {
	case *BoundBinaryExpression:
		return v.Kind().String()
	case *BoundUnaryExpression:
		return v.Kind().String()
	default:
		if v != nil {
			return v.Kind().String()
		} else {
			return "NIL"
		}

	}
}
