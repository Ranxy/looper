package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/evaluator"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
)

func main() {
	showTree := false

	reader := bufio.NewReader(os.Stdin)

	vm := bind.NewVariableManage()

	textBuild := strings.Builder{}

	for {
		if textBuild.Len() == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("| ")
		}

		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		if textBuild.Len() == 0 {
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
		}

		textBuild.Write(line)

		text := textBuild.String()

		sourceText := texts.NewTextSource([]rune(text))

		tree := syntax.NewParser(sourceText).Parse()

		if showTree {
			tree.Print(os.Stdout)
		}
		if len(tree.Diagnostics.List) != 0 {
			tree.Diagnostics.PrintWithSource(sourceText)
		} else {
			b := bind.NewBinder(vm)
			boundExpress := b.BindExpression(tree.Root)
			if len(b.Diagnostics.List) != 0 {
				b.Diagnostics.PrintWithSource(sourceText)
			} else {
				eval := evaluator.NewEvaluater(boundExpress, vm)
				res := eval.Evaluate()
				fmt.Println(res)
			}
		}

		textBuild.Reset()
	}
}
