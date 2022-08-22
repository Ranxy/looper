package optimize

import (
	"container/list"

	"github.com/Ranxy/looper/bind"
)

func FlattenStatement(statement bind.Boundstatement) *bind.BoundBlockStatements {
	ss := make([]bind.Boundstatement, 0)
	stack := list.New()

	stack.PushBack(statement)

	for stack.Len() > 0 {
		current := stack.Back()
		if block, ok := current.Value.(*bind.BoundBlockStatements); ok {
			for i := len(block.Statement) - 1; i >= 0; i-- {
				stack.PushBack(block.Statement[i])
			}
		} else {
			ss = append(ss, current.Value.(bind.Boundstatement))
		}
		stack.Remove(current)
	}
	return bind.NewBoundBlockStatement(ss)
}
