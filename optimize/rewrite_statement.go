package optimize

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
)

type RewriteStatement interface {
	RewriteStatement(w Rewrite, node bind.Boundstatement) bind.Boundstatement
	RewriteBlockStatement(w Rewrite, node *bind.BoundBlockStatements) bind.Boundstatement
	RewriteVariableDeclaration(w Rewrite, node *bind.BoundVariableDeclaration) bind.Boundstatement
	RewriteIfStatement(w Rewrite, node *bind.BoundIfStatements) bind.Boundstatement
	RewriteWhileStatement(w Rewrite, node *bind.BoundWhileStatements) bind.Boundstatement
	RewriteForStatement(w Rewrite, node *bind.BoundForStatements) bind.Boundstatement
	RewriteLabelStatement(w Rewrite, node *bind.LabelStatement) bind.Boundstatement
	RewriteGotoStatement(w Rewrite, node *bind.GotoStatement) bind.Boundstatement
	RewriteCondGotoStatement(w Rewrite, node *bind.ConditionalGotoStatement) bind.Boundstatement
	RewriteExpressionStatement(w Rewrite, node *bind.BoundExpressStatements) bind.Boundstatement
}

func (b *BasicRewrite) RewriteStatement(w Rewrite, node bind.Boundstatement) bind.Boundstatement {
	switch node.Kind() {
	case bind.BoundNodeKindBlockStatement:
		return w.RewriteBlockStatement(w, node.(*bind.BoundBlockStatements))
	case bind.BoundNodeKindVariableDeclaration:
		return w.RewriteVariableDeclaration(w, node.(*bind.BoundVariableDeclaration))
	case bind.BoundNodeKindIfStatement:
		return w.RewriteIfStatement(w, node.(*bind.BoundIfStatements))
	case bind.BoundNodeKindWhileStatement:
		return w.RewriteWhileStatement(w, node.(*bind.BoundWhileStatements))
	case bind.BoundNodeKindForStatement:
		return w.RewriteForStatement(w, node.(*bind.BoundForStatements))
	case bind.BoundNodeKindLabelStatement:
		return w.RewriteLabelStatement(w, node.(*bind.LabelStatement))
	case bind.BoundNodeKindGotoStatement:
		return w.RewriteGotoStatement(w, node.(*bind.GotoStatement))
	case bind.BoundNodeKindConditionalGotoStatement:
		return w.RewriteCondGotoStatement(w, node.(*bind.ConditionalGotoStatement))
	case bind.BoundNodeKindExpressionStatement:
		return w.RewriteExpressionStatement(w, node.(*bind.BoundExpressStatements))
	default:
		panic(fmt.Sprintf("Unexcepted node %v", node.Kind()))
	}
}

func (b *BasicRewrite) RewriteBlockStatement(w Rewrite, node *bind.BoundBlockStatements) bind.Boundstatement {
	list := make([]bind.Boundstatement, 0)
	change := false

	for i := 0; i < len(node.Statement); i++ {
		oldStatement := node.Statement[i]
		newStatement := w.RewriteStatement(w, node.Statement[i])
		if newStatement != oldStatement {
			change = true
			list = append(list, newStatement)
		} else {
			list = append(list, oldStatement)
		}
	}
	if !change {
		return node
	}

	return bind.NewBoundBlockStatement(list)
}

func (b *BasicRewrite) RewriteVariableDeclaration(w Rewrite, node *bind.BoundVariableDeclaration) bind.Boundstatement {
	initializer := w.RewriteExpression(w, node.Initializer)
	if initializer == node.Initializer {
		return node
	}
	return bind.NewBoundVariableDeclaration(node.Variable, initializer)
}

func (b *BasicRewrite) RewriteIfStatement(w Rewrite, node *bind.BoundIfStatements) bind.Boundstatement {
	cond := w.RewriteExpression(w, node.Condition)
	then := w.RewriteStatement(w, node.ThenStatement)
	var elseStatement bind.Boundstatement
	if node.ElseStatement != nil {
		elseStatement = w.RewriteStatement(w, node.ElseStatement)
	}
	if cond == node.Condition && then == node.ThenStatement && elseStatement == node.ElseStatement {
		return node
	}

	return bind.NewBoundIfStatements(cond, then, elseStatement)
}

func (b *BasicRewrite) RewriteWhileStatement(w Rewrite, node *bind.BoundWhileStatements) bind.Boundstatement {
	cond := w.RewriteExpression(w, node.Condition)
	body := w.RewriteStatement(w, node.Body)
	if cond == node.Condition && body == node.Body {
		return node
	}

	return bind.NewBoundWhileStatements(cond, body, node.BreakLabel, node.ContinueLabel)
}

func (b *BasicRewrite) RewriteForStatement(w Rewrite, node *bind.BoundForStatements) bind.Boundstatement {
	init := w.RewriteStatement(w, node.InitCondition)
	end := w.RewriteExpression(w, node.EndCheckConditionExpress)
	update := w.RewriteStatement(w, node.UpdateCondition)
	body := w.RewriteStatement(w, node.Body)
	if init == node.InitCondition && end == node.EndCheckConditionExpress && update == node.UpdateCondition && body == node.Body {
		return node
	}
	return bind.NewBoundForStatements(init, end, update, body, node.BreakLabel, node.ContinueLabel)
}

func (b *BasicRewrite) RewriteLabelStatement(w Rewrite, node *bind.LabelStatement) bind.Boundstatement {
	return node
}

func (b *BasicRewrite) RewriteGotoStatement(w Rewrite, node *bind.GotoStatement) bind.Boundstatement {
	return node
}

func (b *BasicRewrite) RewriteCondGotoStatement(w Rewrite, node *bind.ConditionalGotoStatement) bind.Boundstatement {
	cond := w.RewriteExpression(w, node.Condition)
	if cond == node.Condition {
		return node
	}
	return bind.NewConditionalGotoSymbol(node.Label, cond, node.JumpIfFalse)
}
func (b *BasicRewrite) RewriteExpressionStatement(w Rewrite, node *bind.BoundExpressStatements) bind.Boundstatement {
	express := w.RewriteExpression(w, node.Express)
	if express == node.Express {
		return node
	}
	return bind.NewBoundExpressStatements(express)
}
