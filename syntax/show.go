package syntax

import (
	"fmt"
	"io"
)

func PrintExpress(writer io.Writer, express Express, indent string, isLast bool) {

	var marker string
	if isLast {
		marker = "└──"
	} else {
		marker = "├──"
	}

	writer.Write([]byte(indent))
	writer.Write([]byte(marker))

	writer.Write([]byte(express.Kind().String()))
	if token, ok := express.(SyntaxToken); ok && token.Value != nil {
		writer.Write([]byte(" "))
		writer.Write([]byte(fmt.Sprint(token.Value)))
	}
	writer.Write([]byte("\n"))

	if isLast {
		indent += "    "
	} else {
		indent += "│   "
	}

	for idx, child := range express.GetChildren() {
		isLast := idx == len(express.GetChildren())-1
		PrintExpress(writer, child, indent, isLast)
	}
}
