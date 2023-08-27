package texts_test

import (
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/bind/program"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestTextSource(t *testing.T) {
	type args struct {
		text []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				text: []rune("."),
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				text: []rune(".\r\n"),
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				text: []rune(".\r\n\r\n"),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := texts.NewTextSource(tt.args.text)
			if len(got.Lines) != tt.want {
				t.Errorf("TextSource Len = %v, want %v", len(got.Lines), tt.want)
			}
		})
	}
}

func Test_TextSource(t *testing.T) {
	text := `
	fn doSomething(x:int, f:string):int{
		for var i = 0;i<x;i=i+1{
			print(f)
		} 
		return x
	}
	let x:int = "a"
	fn hello(v:int){
		return ()
	}
	{
		return 2
	}
	
	hello(doSomething(2,"abc"))
	`

	source := texts.NewTextSource([]rune(text))

	require.Len(t, source.Lines, 17)

	tree := syntax.ParseToTree(source)
	boundTree := bind.BindGlobalScope(nil, tree.Root)
	program := program.BindProgram(boundTree)
	require.True(t, program.Diagnostic.Has())

	program.Diagnostic.PrintWithSourceStdout(source)

	require.Len(t, program.Statement.Statement, 3)

}
