package evaluator

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/syntax"
)

type Evaluater struct {
	root bind.Boundstatement
	vm   map[syntax.VariableSymbol]any

	lastValue any
}

func NewEvaluater(root bind.BoundExpression, vm map[syntax.VariableSymbol]any) *Evaluater {
	return &Evaluater{
		root: root,
		vm:   vm,
	}
}

func (e *Evaluater) Evaluate() any {
	e.EvaluateStatement(e.root)
	return e.lastValue
}

func (e *Evaluater) EvaluateStatement(node bind.Boundstatement) {
	switch node.Kind() {
	case bind.BoundNodeKindBlockStatement:
		e.EvaluateBlockStatement(node.(*bind.BoundBlockStatements))
	case bind.BoundNodeKindVariableDeclaration:
		e.EvaluateVariableDeclaration(node.(*bind.BoundVariableDeclaration))
	case bind.BoundNodeKindIfStatement:
		e.EvaluateIfStatement(node.(*bind.BoundIfStatements))
	case bind.BoundNodeKindWhileStatement:
		e.EvaluateWhileStatement(node.(*bind.BoundWhileStatements))
	case bind.BoundNodeKindForStatement:
		e.EvaluateForStatement(node.(*bind.BoundForStatements))
	case bind.BoundNodeKindExpressionStatement:
		e.EvaluateExpressionStatement(node.(*bind.BoundExpressStatements))
	default:
		panic(fmt.Sprintf("Unexceped Node %v", node.Kind()))
	}
}

func (e *Evaluater) EvaluateForStatement(node *bind.BoundForStatements) {
	e.EvaluateStatement(node.InitCondition)
	for e.EvaluateExpression(node.EndCondition).(bool) {
		e.EvaluateStatement(node.Body)
		e.EvaluateStatement(node.UpdateCondition)
	}
}

func (e *Evaluater) EvaluateWhileStatement(node *bind.BoundWhileStatements) {
	for e.EvaluateExpression(node.Condition).(bool) {
		e.EvaluateStatement(node.Body)
	}
}

func (e *Evaluater) EvaluateIfStatement(node *bind.BoundIfStatements) {
	cond := e.EvaluateExpression(node.Condition).(bool)
	if cond {
		e.EvaluateStatement(node.ThenStatement)
	} else if node.ElseStatement != nil {
		e.EvaluateStatement(node.ElseStatement)
	} else {
		e.lastValue = 0 //todo use unit type
	}
}

func (e *Evaluater) EvaluateBlockStatement(node *bind.BoundBlockStatements) {
	for _, statement := range node.Statement {
		e.EvaluateStatement(statement)
	}
}

func (e *Evaluater) EvaluateVariableDeclaration(node *bind.BoundVariableDeclaration) {
	value := e.EvaluateExpression(node.Initializer)
	e.vm[*node.Variable] = value
	e.lastValue = value
}

func (e *Evaluater) EvaluateExpressionStatement(node *bind.BoundExpressStatements) {
	e.lastValue = e.EvaluateExpression(node.Express)
}

func (e *Evaluater) EvaluateExpression(node bind.BoundExpression) any {

	switch n := node.(type) {
	case *bind.BoundLiteralExpression:
		return e.evaluateLiteralExpression(n)
	case *bind.BoundVariableExpression:
		return e.evaluateVariableExpression(n)
	case *bind.BoundAssignmentExpression:
		return e.evaluateAssignmentExpression(n)
	case *bind.BoundUnaryExpression:
		return e.evaluateUnaryExpression(n)
	case *bind.BoundBinaryExpression:
		return e.evaluateBinaryExpression(n)
	default:
		panic(fmt.Sprintf("Unexceped node %v", node))
	}
}

func (e *Evaluater) evaluateLiteralExpression(node *bind.BoundLiteralExpression) any {
	return node.Value
}

func (e *Evaluater) evaluateVariableExpression(node *bind.BoundVariableExpression) any {
	v := e.vm[*node.Variable]
	if v == nil {
		panic(fmt.Sprintf("Undefined Variable %s", node.Variable.Name))
	}
	return v
}

func (e *Evaluater) evaluateAssignmentExpression(node *bind.BoundAssignmentExpression) any {
	value := e.EvaluateExpression(node.Express)
	e.vm[*node.Variable] = value
	return value
}

func (e *Evaluater) evaluateUnaryExpression(node *bind.BoundUnaryExpression) any {
	operand := e.EvaluateExpression(node.Operand)
	switch node.Op.Kind {
	case bind.BoundUnaryOperatorKindIdentity:
		return operand.(int64)
	case bind.BoundUnaryOperatorKindNegation:
		return -(operand.(int64))
	case bind.BoundUnaryOperatorKindLogicalNegation:
		return !(operand.(bool))
	default:
		panic(fmt.Sprintf("Unexceped unary operator:%v", node.Op))
	}
}

func (e *Evaluater) evaluateBinaryExpression(node *bind.BoundBinaryExpression) any {
	left := e.EvaluateExpression(node.Left)
	right := e.EvaluateExpression(node.Right)

	switch node.Op.Kind {
	case bind.BoundBinaryKindAddition:
		return left.(int64) + right.(int64)
	case bind.BoundBinaryKindSubtraction:
		return left.(int64) - right.(int64)
	case bind.BoundBinaryKindMultiplication:
		return left.(int64) * right.(int64)
	case bind.BoundBinaryKindDivision:
		return left.(int64) / right.(int64)
	case bind.BoundBinaryKindLogicalAnd:
		return left.(bool) && right.(bool)
	case bind.BoundBinaryKindLogicalOr:
		return left.(bool) || right.(bool)
	case bind.BoundBinaryKindEquals:
		return left == right
	case bind.BoundBinaryKindNotEquals:
		return left != right
	case bind.BoundBinaryKindLess:
		return left.(int64) < right.(int64)
	case bind.BoundBinaryKindLessEqual:
		return left.(int64) <= right.(int64)
	case bind.BoundBinaryKindGreat:
		return left.(int64) > right.(int64)
	case bind.BoundBinaryKindGreatEqual:
		return left.(int64) >= right.(int64)

	default:
		panic(fmt.Sprintf("Unexceped binary operator:%q and", node.Op))
	}
}
