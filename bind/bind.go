package bind

import (
	"fmt"

	"github.com/Ranxy/looper/syntax"
)

type Binder struct {
	Errors []string
	vm     VariableManage
}

func NewBinder(vm VariableManage) *Binder {
	return &Binder{
		Errors: make([]string, 0),
		vm:     vm,
	}
}

func (b *Binder) BindExpression(express syntax.Express) BoundExpression {
	switch express.Kind() {
	case syntax.SyntaxKindLiteralExpress:
		return b.BindLiteralExpress(express.(*syntax.LiteralExpress))
	case syntax.SyntaxKindNameExpress:
		return b.BindNameExpress(express.(*syntax.NameExpress))
	case syntax.SyntaxKindUnaryExpress:
		return b.BindUnaryExpress(express.(*syntax.UnaryExpress))
	case syntax.SyntaxKindBinaryExpress:
		return b.BindBinaryOperator(express.(*syntax.BinaryExpress))
	case syntax.SyntaxKindParenthesizedExpress:
		return b.BindExpression(express.(*syntax.ParenthesisExpress).Expr)
	case syntax.SyntaxKindAssignmentExpress:
		return b.BindAssignmentExpress(express.(*syntax.AssignmentExpress))
	default:
		panic(fmt.Sprintf("unexceped expresss %q", express.Kind()))
	}
}

func (b *Binder) BindLiteralExpress(express *syntax.LiteralExpress) BoundExpression {
	value := express.Value
	if value == nil {
		value = int64(0)
	}
	return NewBoundLiteralExpression(value)
}

func (b *Binder) BindNameExpress(express *syntax.NameExpress) BoundExpression {
	name := express.Identifier.Text
	variable := b.vm.GetSymbol(name)
	if variable == nil {
		b.Errors = append(b.Errors, fmt.Sprintf("Undefined variable '%s'", name))
		return NewBoundLiteralExpression(int64(0))
	}
	return NewBoundVariableExpression(variable)
}

func (b *Binder) BindAssignmentExpress(express *syntax.AssignmentExpress) BoundExpression {

	name := express.Identifier.Text
	boundExpress := b.BindExpression(express.Express)

	variable := syntax.NewVariableSymbol(name, boundExpress.Type())
	b.vm.Declare(variable)
	return NewBoundAssignmentExpression(variable, boundExpress)
}

func (b *Binder) BindUnaryExpress(express *syntax.UnaryExpress) BoundExpression {
	boundOperand := b.BindExpression(express.Operand)
	boundOperator := BindBoundUnaryOperator(express.Operator.Kind(), boundOperand.Type())

	if boundOperator == nil {
		b.Errors = append(b.Errors, fmt.Sprintf("Unary operator %s is not defined for type %q", express.Operator.Text, boundOperand.Type()))
		return boundOperand
	}

	return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder) BindBinaryOperator(express *syntax.BinaryExpress) BoundExpression {
	boundLeft := b.BindExpression(express.Left)
	boundRight := b.BindExpression(express.Right)
	boundOperator := BindBoundBinaryOperator(express.Operator.Kind(), boundLeft.Type(), boundRight.Type())

	if boundOperator == nil {
		b.Errors = append(b.Errors, fmt.Sprintf("Binary operator %s is not defined for type %q and type %q", express.Operator.Text, boundLeft.Type(), boundRight.Type()))
		return boundLeft
	}

	return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}
