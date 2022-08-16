package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ranxy/looper/compilation"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
)

func main() {
	showTree := false

	reader := bufio.NewReader(os.Stdin)

	vm := make(map[syntax.VariableSymbol]any)
	var previous *compilation.Compilation

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
			if text == "#reset" {
				previous = nil
				vm = make(map[syntax.VariableSymbol]any)
				fmt.Println("Success")
				continue
			}
			if text == "#dump" {
				fmt.Print(vm)
				continue
			}
			if text == "" {
				continue
			}
		}

		textBuild.Write(line)

		text := textBuild.String()

		sourceText := texts.NewTextSource([]rune(text))

		tree := syntax.ParseToTree(sourceText)

		if showTree {
			err = tree.Print(os.Stdout)
			if err != nil {
				panic(fmt.Sprintf("ShowTreeFailed %v", err))
			}
		}
		if len(tree.Diagnostics.List) != 0 {
			tree.Diagnostics.PrintWithSource(sourceText)
			tree.Diagnostics.Reset()
		} else {
			cm := compilation.NewCompliation(previous, tree)

			res := cm.Evaluate(vm)
			if res.Diagnostic.Has() {
				res.Diagnostic.PrintWithSource(sourceText)
				res.Diagnostic.Reset()
			} else {
				fmt.Println(res.Value)
			}
			previous = cm
		}

		textBuild.Reset()
	}
}
