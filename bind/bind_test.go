package bind_test

import (
	"os"
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/bind/program"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestCase(t *testing.T) {
	tests := []struct {
		text string
	}{
		{
			text: "1 + 2",
		},
		{
			text: "10 - 2 * 3",
		},
		{
			text: "1 + 2 * (3 -1 )",
		},
		{
			text: "-2",
		},
		{
			text: "1+-2",
		},
		{
			text: "-(1 * 2) -3",
		},
		{
			text: "1 < 2",
		},
		{
			text: "1 <= 2",
		},
		{
			text: "1 > 2",
		},
		{
			text: "1 >= 2",
		},
		{
			text: "1 & 2",
		},
		{
			text: "1 | 2",
		},
		{
			text: "1 ^ 2",
		},
		{
			text: "~1",
		},
		{
			text: "1 & 2 | 3 ^ ~ 4",
		},
		{
			text: `"hello" + "world"`,
		},
		{
			text: `{var a = "hello" a+ "world"}`,
		},
		{
			text: `{var a = "" a+ "foo"}`,
		},
		{
			text: "{ var a = 0 if a == 2 a = 3 else a = 6 a }",
		},
		{
			text: "{ var a = 7 if a == 2 a = 3 a }",
		},
		{
			text: "{ var i = 10 var result = 0 while i != 0 { result = result + i i = i - 1} result }",
		},
		{
			text: "{ var result = 0 for var i = 1; i < 5; i = i + 1 { result = result + i } result }",
		},
		{
			text: "{var i = 0 var result = 0 for i = 1; i < 5; i = i + 1 { result = result + i } result }",
		},
		{
			text: `{print("hello"+"world")}`,
		},
		{
			text: `{let a = randint(2+2)}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.ParseToTree(textSource)
			boundTree := bind.BindGlobalScope(nil, tree.Root)
			if boundTree.Diagnostic.Has() {
				boundTree.Diagnostic.Print(textSource)
			}
			require.False(t, boundTree.Diagnostic.Has())
		})
	}
}

func TestBinder_BindExpression(t *testing.T) {
	text := "let a = 2"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)
	require.Len(t, boundTree.Variables, 1)
	require.Equal(t, boundTree.Variables[0].GetName(), "a")
	require.Equal(t, boundTree.Variables[0].IsReadOnly(), true)
}

func TestBinder_BindifStatement(t *testing.T) {
	text := "{var b = 0{let a = 2+1 { if(-a == 1){ b = 2}else{ b = 3}}}}"
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)

	program := program.BindProgram(boundTree)

	err := bind.PrintBoundProgram(os.Stdout, program.Functions, program.Statement)
	require.NoError(t, err)

	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)

	require.False(t, boundTree.Diagnostic.Has())
}

func TestBinder_BindWhileStatement(t *testing.T) {
	text := `{ 
		var i = 10 
		var result = 0 
		while i != 0 {
			if (i == 5){
				continue
			}
			if (i == 9){
				break
			}
			result = result + i 
			if (i == 0){
				break
			}
			i = i - 1
		} 
		result
	}`
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	program := program.BindProgram(boundTree)
	err := bind.PrintBoundProgram(os.Stdout, program.Functions, program.Statement)
	require.NoError(t, err)
	t.Log(boundTree.Diagnostic)
	t.Log(boundTree.Variables)

	require.False(t, boundTree.Diagnostic.Has())
}

func TestBinder_BindFunction(t *testing.T) {
	text := `fn doSomething(x:int, f:string):int{for var i = 0;i<x;i=i+1{print(f)} return x} {doSomething(3,"abc") }`
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	program := program.BindProgram(boundTree)
	err := bind.PrintBoundProgram(os.Stdout, program.Functions, program.Statement)
	require.NoError(t, err)
	require.Equal(t, len(program.Functions), 1)

	if program.Diagnostic.Has() {
		program.Diagnostic.Print(textSource)
	}
}

func TestBinder_BindFunction_multiple(t *testing.T) {
	text := `
	fn doSomething(x:int, f:string):int{
		for var i = 0;i<x;i=i+1{
			print(f)
		} 
		return x
	}
	fn hello(v:int){
		return ()
	}
	
	hello(doSomething(2,"abc"))
	`
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	program := program.BindProgram(boundTree)
	err := bind.PrintBoundProgram(os.Stdout, program.Functions, program.Statement)
	require.NoError(t, err)

	require.Equal(t, len(program.Functions), 2)

	if program.Diagnostic.Has() {
		program.Diagnostic.Print(textSource)
	}
}

func TestBinder_Return_Must_In_function(t *testing.T) {
	text := `
	var i = 10 
	var result = 0 
	{
		return 3
	}
	`

	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	program := program.BindProgram(boundTree)
	err := bind.PrintBoundProgram(os.Stdout, program.Functions, program.Statement)
	require.NoError(t, err)

	require.True(t, program.Diagnostic.Has())

	require.Len(t, program.Diagnostic.List, 1)
	require.Equal(t, "The keyword return can only be used inside of function", program.Diagnostic.List[0].Message)
}
