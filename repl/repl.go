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

		tree := syntax.NewParser(text).Parse()

		if showTree {
			tree.Print()
		}
		if len(tree.Errors) != 0 {
			for _, err := range tree.Errors {
				fmt.Println(err)
			}
		} else {
			b := bind.NewBinder(vm)
			boundExpress := b.BindExpression(tree.Root)
			if len(b.Errors) != 0 {
				for _, err := range b.Errors {
					fmt.Println(err)
				}
				continue
			}
			eval := evaluator.NewEvaluater(boundExpress, vm)
			res := eval.Evaluate()
			fmt.Println(res)
		}

	}
}
