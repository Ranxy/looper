package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/evaluator"
	"github.com/Ranxy/looper/syntax"
)

func main() {
	showTree := false

	reader := bufio.NewReader(os.Stdin)

	vm := bind.NewVariableManage()

	for {
		fmt.Print("> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		text := string(line)

		if text == "#showTree" {
			showTree = !showTree
			if showTree {
				fmt.Println("Showing parse tree")
			} else {
				fmt.Println("Not showing parse tree")
			}
			continue
		}
		if text == "#dump" {
			fmt.Print(vm.Dump())
			continue
		}
		if text == "" {
			continue
		}

		tree := syntax.NewParser(text).Parse()

		if showTree {
			tree.Print()
		}
		if len(tree.Diagnostics.List) != 0 {
			tree.Diagnostics.Print(text)
		} else {
			b := bind.NewBinder(vm)
			boundExpress := b.BindExpression(tree.Root)
			if len(b.Diagnostics.List) != 0 {
				b.Diagnostics.Print(text)
			} else {
				eval := evaluator.NewEvaluater(boundExpress, vm)
				res := eval.Evaluate()
				fmt.Println(res)
			}
		}

	}
}
