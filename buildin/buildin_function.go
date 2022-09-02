package buildin

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/Ranxy/looper/symbol"
)

var AllBuildinFunc = []*symbol.FunctionSymbol{FunctionPrint, FunctionInputStr, FunctionRnd}

var (
	FunctionPrint    = symbol.NewFunctionSymbol("print", []*symbol.ParameterSymbol{symbol.NewParameterSymbol("text", symbol.TypeString)}, symbol.TypeUnit)
	FunctionInputStr = symbol.NewFunctionSymbol("inputstr", []*symbol.ParameterSymbol{}, symbol.TypeString)
	FunctionRnd      = symbol.NewFunctionSymbol("randint", []*symbol.ParameterSymbol{symbol.NewParameterSymbol("max", symbol.TypeInt)}, symbol.TypeInt)
)

func FunctionPrintImpl(text string) {
	fmt.Println(text)
}

func FunctionInputStrImpl() string {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func FunctionRndImpl(max int64) int64 {
	return rand.Int63n(max)
}
