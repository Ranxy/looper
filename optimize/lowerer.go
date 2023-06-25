package optimize

import (
	"fmt"

	"github.com/Ranxy/looper/bind"
)

var _ Rewrite = &LowerRewrite{}

type LowerRewrite struct {
	Rewrite
	labelCount int64
}

func Lower(statement bind.Boundstatement) *bind.BoundBlockStatements {
	l := &LowerRewrite{&BasicRewrite{}, 1}
	res := l.RewriteStatement(l, statement)

	return FlattenStatement(res)
}

func (l *LowerRewrite) GenerateLabel() *bind.BoundLabel {
	name := fmt.Sprintf("Label%d", l.labelCount)
	l.labelCount += 1
	return bind.NewBoundLabel(name)
}

func (l *LowerRewrite) RewriteIfStatement(w Rewrite, node *bind.BoundIfStatements) bind.Boundstatement {
	if node.ElseStatement == nil {
		endLabel := l.GenerateLabel()
		// if false, goto endlabel
		gotoFalse := bind.NewConditionalGotoSymbol(endLabel, node.Condition, true)
		endLabelStatement := bind.NewLabelSymbol(endLabel)

		res := bind.NewBoundBlockStatement([]bind.Boundstatement{gotoFalse, node.ThenStatement, endLabelStatement})
		return w.RewriteStatement(w, res)
	} else {

		/*
			     ┌──────── gotoFalse
			     │             │
			if false           │
			goto the           ▼
			else label     ThenStatement
			     │             │
			     │             │
			     │             ▼
			     │         gotoEnd  ──────────────────┐
			     │                                    │
			     │                                    │
			     │                           if then statement
			     └────────►elseLabel         done, goto endLabel
			                   │                      │
			                   │                      │
			                   ▼                      │
			               ElseStatement              │
			                   │                      │
			                   │                      │
			                   ▼                      │
			               endLabel ◄─────────────────┘
		*/
		elseLabel := l.GenerateLabel()
		endLabel := l.GenerateLabel()
		gotoFalse := bind.NewConditionalGotoSymbol(elseLabel, node.Condition, true)
		gotoEnd := bind.NewGotoSymbol(endLabel)
		elseLabelStatement := bind.NewLabelSymbol(elseLabel)
		endLabelStatement := bind.NewLabelSymbol(endLabel)
		res := bind.NewBoundBlockStatement([]bind.Boundstatement{gotoFalse, node.ThenStatement, gotoEnd, elseLabelStatement, node.ElseStatement, endLabelStatement})

		return w.RewriteStatement(w, res)
	}
}

func (l *LowerRewrite) RewriteWhileStatement(w Rewrite, node *bind.BoundWhileStatements) bind.Boundstatement {

	// goto check
	// continue:
	// <body>
	// check:
	// gotoTrue <condition> continue
	// break:

	checkLabel := l.GenerateLabel()
	gotoCheck := bind.NewGotoSymbol(checkLabel)

	continueLabelStatement := bind.NewLabelSymbol(node.ContinueLabel)
	checklabelStatement := bind.NewLabelSymbol(checkLabel)
	gotoTrue := bind.NewConditionalGotoSymbol(node.ContinueLabel, node.Condition, true)
	breakLabelStatement := bind.NewLabelSymbol(node.BreakLabel)

	result := bind.NewBoundBlockStatement([]bind.Boundstatement{
		gotoCheck,
		continueLabelStatement,
		node.Body,
		checklabelStatement,
		gotoTrue,
		breakLabelStatement,
	})
	return w.RewriteStatement(w, result)
}

func (l *LowerRewrite) RewriteForStatement(w Rewrite, node *bind.BoundForStatements) bind.Boundstatement {
	// todo
	labelStart := l.GenerateLabel()
	labelEnd := l.GenerateLabel()
	condJumpEnd := bind.NewConditionalGotoSymbol(labelEnd, node.EndCheckConditionExpress, true)
	jumpStart := bind.NewGotoSymbol(labelStart)
	startStatement := bind.NewLabelSymbol(labelStart)
	endStatement := bind.NewLabelSymbol(labelEnd)

	res := bind.NewBoundBlockStatement([]bind.Boundstatement{
		node.InitCondition,
		startStatement,
		condJumpEnd,
		node.Body,
		node.UpdateCondition,
		jumpStart,
		endStatement,
	})

	return w.RewriteStatement(w, res)
}
