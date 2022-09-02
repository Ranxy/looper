package evaluator

import (
	"fmt"
	"os"
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/optimize"
	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/texts"
	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		text    string
		want    any
		wantErr bool
	}{
		{
			text:    "1 + 2",
			want:    int64(3),
			wantErr: false,
		},
		{
			text:    "10 - 2 * 3",
			want:    int64(4),
			wantErr: false,
		},
		{
			text:    "1 + 2 * (3 -1 )",
			want:    int64(5),
			wantErr: false,
		},
		{
			text:    "-2",
			want:    int64(-2),
			wantErr: false,
		},
		{
			text:    "1+-2",
			want:    int64(-1),
			wantErr: false,
		},
		{
			text:    "-(1 * 2) -3",
			want:    int64(-5),
			wantErr: false,
		},
		{
			text:    "1 < 2",
			want:    true,
			wantErr: false,
		},
		{
			text:    "1 <= 2",
			want:    true,
			wantErr: false,
		},
		{
			text:    "2 <= 2",
			want:    true,
			wantErr: false,
		},
		{
			text:    "3 < 2",
			want:    false,
			wantErr: false,
		},
		{
			text:    "1 > 2",
			want:    false,
			wantErr: false,
		},
		{
			text:    "2 > 2",
			want:    false,
			wantErr: false,
		},
		{
			text:    "2 >= 2",
			want:    true,
			wantErr: false,
		},
		{
			text:    "3 > 2",
			want:    true,
			wantErr: false,
		},
		{
			text: "~1",
			want: int64(^1),
		},
		{
			text: "1 & 2",
			want: int64(1 & 2),
		},
		{
			text: "1 | 2",
			want: int64(1 | 2),
		},
		{
			text: "1 ^ 2",
			want: int64(1 ^ 2),
		},
		{
			text: "1 & 2 | 3 ^ ~ 4",
			want: int64(1&2 | 3 ^ ^4),
		},
		{
			text:    "{ var a = 0 if a == 2 a = 3 else a = 6 a }",
			want:    int64(6),
			wantErr: false,
		},
		{
			text:    "{ var a = 7 if a == 2 a = 3 a }",
			want:    int64(7),
			wantErr: false,
		},
		{
			text: "{ var i = 10 var result = 0 while i != 0 { result = result + i i = i - 1} result }",
			want: int64(55),
		},
		{
			text: "{ var result = 0 for var i = 1; i < 5; i = i + 1 { result = result + 1 } result }",
			want: int64(4),
		},
		{
			text: "{var i = 0 var result = 0 for i = 1; i < 5 ;i = i + 1 { result = result + 1 } result }",
			want: int64(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.ParseToTree(textSource)

			boundTree := bind.BindGlobalScope(nil, tree.Root)
			if len(boundTree.Diagnostic.List) != 0 {
				boundTree.Diagnostic.Print(tt.text)
				t.FailNow()
			}
			vm := make(map[symbol.VariableSymbol]any)

			blockStatements := optimize.Lower(boundTree.Statements)
			ev := NewEvaluater(blockStatements, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}
func TestEvaluate_bool(t *testing.T) {
	tests := []struct {
		text    string
		want    any
		wantErr bool
	}{
		{
			text:    "true && false",
			want:    false,
			wantErr: false,
		},
		{
			text:    "true == false",
			want:    false,
			wantErr: false,
		},
		{
			text:    "true == true",
			want:    true,
			wantErr: false,
		},
		{
			text:    "true || true ==false",
			want:    true,
			wantErr: false,
		},
		{
			text:    "true && true ==(false && true)",
			want:    false,
			wantErr: false,
		},
		{
			text:    "true && 1==2",
			want:    false,
			wantErr: false,
		},
		{
			text:    "true && 1==1",
			want:    true,
			wantErr: false,
		},
		{
			text:    "!false && 3 == 7- 2*2",
			want:    true,
			wantErr: false,
		},
		{
			text:    "!false && !(3!=3) && !(3 == 7 - 2*2)",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.ParseToTree(textSource)
			boundTree := bind.BindGlobalScope(nil, tree.Root)
			if len(boundTree.Diagnostic.List) != 0 {
				boundTree.Diagnostic.Print(tt.text)
				t.FailNow()
			}
			vm := make(map[symbol.VariableSymbol]any)

			blockStatements := optimize.Lower(boundTree.Statements)
			ev := NewEvaluater(blockStatements, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEvaluate_string(t *testing.T) {
	tests := []struct {
		text    string
		want    any
		wantErr bool
	}{
		{
			text:    `"hello" + " "+ "world"`,
			want:    "hello world",
			wantErr: false,
		},
		{
			text:    `{var a = "" a+ "foo"}`,
			want:    "foo",
			wantErr: false,
		},
		{
			text:    `{ var res = 0 var a = "foo" if a =="foo"{res = 1}else{res = 2} res}`,
			want:    int64(1),
			wantErr: false,
		},
		{
			text:    `{var res = "" {for var i = 0;i<10; i=i+1{res = res + "f"}} res}`,
			want:    "ffffffffff",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.ParseToTree(textSource)
			boundTree := bind.BindGlobalScope(nil, tree.Root)
			if len(boundTree.Diagnostic.List) != 0 {
				boundTree.Diagnostic.Print(tt.text)
				t.FailNow()
			}
			vm := make(map[symbol.VariableSymbol]any)

			blockStatements := optimize.Lower(boundTree.Statements)

			// bind.PrintBoundTree(os.Stdout, blockStatements)
			ev := NewEvaluater(blockStatements, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEvaluate_function(t *testing.T) {
	tests := []struct {
		text    string
		want    any
		wantErr bool
	}{
		{
			text:    `{let a = randint(100)}`,
			want:    int64(10),
			wantErr: false,
		},
		{
			text:    `{let a = "hello" { print(a + "world")}}`,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.ParseToTree(textSource)
			boundTree := bind.BindGlobalScope(nil, tree.Root)
			if len(boundTree.Diagnostic.List) != 0 {
				boundTree.Diagnostic.Print(tt.text)
				t.FailNow()
			}
			vm := make(map[symbol.VariableSymbol]any)

			blockStatements := optimize.Lower(boundTree.Statements)

			bind.PrintBoundTree(os.Stdout, blockStatements)
			ev := NewEvaluater(blockStatements, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEvaluate_variable(t *testing.T) {
	vm := make(map[symbol.VariableSymbol]any)

	var bt *bind.BoundGlobalScope

	bt = ev_variable(bt, vm, t, "var a = 1+1", int64(2))
	bt = ev_variable(bt, vm, t, "a=3", int64(3))
	bt = ev_variable(bt, vm, t, "-a", int64(-3))
	bt = ev_variable(bt, vm, t, "a+2", int64(5))

	bt = ev_variable(bt, vm, t, "a==3", true)

	bt = ev_variable(bt, vm, t, "a=a+a+1", int64(7))
	bt = ev_variable(bt, vm, t, "var b = false", false)
	bt = ev_variable(bt, vm, t, "b=(a==7)", true)
	bt = ev_variable(bt, vm, t, "var c = 5", int64(5))
	bt = ev_variable(bt, vm, t, "{if(a==7){c = 2}else{c = 3} c}", int64(2))
	bt = ev_variable(bt, vm, t, "c == 3", false)
	bt = ev_variable(bt, vm, t, "{if(c == 3){c = 10} c}", int64(2))
	bt = ev_variable(bt, vm, t, "c == 2", true)
	_ = ev_variable(bt, vm, t, "b==false", false)

}

func ev_variable(previous *bind.BoundGlobalScope, vm map[symbol.VariableSymbol]any, t *testing.T, text string, want any) *bind.BoundGlobalScope {
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(previous, tree.Root)
	if len(boundTree.Diagnostic.List) != 0 {
		boundTree.Diagnostic.Print(text)
		t.FailNow()
	}
	blockStatements := optimize.Lower(boundTree.Statements)
	ev := NewEvaluater(blockStatements, vm)

	got := ev.Evaluate()
	require.Equal(t, want, got, fmt.Sprintf("Text %s failed", text))

	return boundTree
}
