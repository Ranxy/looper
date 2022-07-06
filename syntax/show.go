package syntax

import "fmt"

func PrintExpress(express Express, indent string, isLast bool) {

	var marker string
	if isLast {
		marker = "└──"
	} else {
		marker = "├──"
	}

	fmt.Print(indent)
	fmt.Print(marker)

	fmt.Print(express.Kind())
	if token, ok := express.(SyntaxToken); ok && token.Value != nil {
		fmt.Print(" ")
		fmt.Print(token.Value)
	}
	fmt.Println()

	if isLast {
		indent += "    "
	} else {
		indent += "│   "
	}

	for idx, child := range express.GetChildren() {
		isLast := idx == len(express.GetChildren())-1
		PrintExpress(child, indent, isLast)
	}
}
