package syntax

import (
	"fmt"
	"io"
)

func PrintExpress(writer io.Writer, express Express, indent string, isLast bool) (err error) {

	var marker string
	if isLast {
		marker = "└──"
	} else {
		marker = "├──"
	}

	_, err = writer.Write([]byte(indent))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(marker))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(express.Kind().String()))
	if err != nil {
		return err
	}
	if token, ok := express.(SyntaxToken); ok && token.Value != nil {
		_, err = writer.Write(append([]byte(" "), []byte(fmt.Sprint(token.Value))...))
		if err != nil {
			return err
		}
	}
	_, err = writer.Write([]byte("\n"))
	if err != nil {
		return err
	}

	if isLast {
		indent += "    "
	} else {
		indent += "│   "
	}

	for idx, child := range express.GetChildren() {
		isLast := idx == len(express.GetChildren())-1
		err = PrintExpress(writer, child, indent, isLast)
		if err != nil {
			return err
		}
	}
	return nil
}
