package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ranxy/looper/syntax"
)

func main() {
	showTree := false

	reader := bufio.NewReader(os.Stdin)

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

		tree := syntax.NewParser(text).Parse()

		if showTree {
			tree.Print()
		}
		if len(tree.Errors) != 0 {
			for _, err := range tree.Errors {
				fmt.Println(err)
			}
		} else {
			res, err := tree.Eval()
			if err != nil {
				fmt.Println("EvalFailed: ", err)
			} else {
				fmt.Println(res)
			}
		}

	}
}
