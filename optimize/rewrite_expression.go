package optimize

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
)

type BasicRewrite struct {
}

type RewriteExpression interface {
	RewriteExpression(w Rewrite, node bind.BoundExpression) bind.BoundExpression
	RewriteLiteralExpression(w Rewrite, node *bind.BoundLiteralExpression) bind.BoundExpression
	RewriteVariableExpression(w Rewrite, node *bind.BoundVariableExpression) bind.BoundExpression
	RewriteAssignmentExpression(w Rewrite, node *bind.BoundAssignmentExpression) bind.BoundExpression
	RewriteUnaryExpression(w Rewrite, node *bind.BoundUnaryExpression) bind.BoundExpression
	RewriteBinaryExpression(w Rewrite, node *bind.BoundBinaryExpression) bind.BoundExpression
}

func (b *BasicRewrite) RewriteExpression(w Rewrite, node bind.BoundExpression) bind.BoundExpression {
	switch node.Kind() {
	case bind.BoundNodeKindLiteralExpress:
		return w.RewriteLiteralExpression(w, node.(*bind.BoundLiteralExpression))
	case bind.BoundNodeKindVariableExpress:
		return w.RewriteVariableExpression(w, node.(*bind.BoundVariableExpression))
	case bind.BoundNodeKindAssignmentExpress:
		return w.RewriteAssignmentExpression(w, node.(*bind.BoundAssignmentExpression))
	case bind.BoundNodeKindUnaryExpress:
		return w.RewriteUnaryExpression(w, node.(*bind.BoundUnaryExpression))
	case bind.BoundNodeKindBinaryExpress:
		return w.RewriteBinaryExpression(w, node.(*bind.BoundBinaryExpression))
	default:
		panic(fmt.Sprintf("Unexcepted node %s", node.Kind()))
	}
}

func (b *BasicRewrite) RewriteLiteralExpression(w Rewrite, node *bind.BoundLiteralExpression) bind.BoundExpression {
	return node
}

func (b *BasicRewrite) RewriteVariableExpression(w Rewrite, node *bind.BoundVariableExpression) bind.BoundExpression {
	return node
}

func (b *BasicRewrite) RewriteAssignmentExpression(w Rewrite, node *bind.BoundAssignmentExpression) bind.BoundExpression {
	express := w.RewriteExpression(w, node.Express)
	if express == node.Express {
		return node
	}
	return bind.NewBoundAssignmentExpression(node.Variable, express)
}

func (b *BasicRewrite) RewriteUnaryExpression(w Rewrite, node *bind.BoundUnaryExpression) bind.BoundExpression {
	operand := w.RewriteExpression(w, node.Operand)
	if operand == node.Operand {
		return node
	}

	return bind.NewBoundUnaryExpression(node.Op, operand)
}

func (b *BasicRewrite) RewriteBinaryExpression(w Rewrite, node *bind.BoundBinaryExpression) bind.BoundExpression {
	left := w.RewriteExpression(w, node.Left)
	right := w.RewriteExpression(w, node.Right)
	if left == node.Left && right == node.Right {
		return node
	}
	return bind.NewBoundBinaryExpression(left, node.Op, right)
}
