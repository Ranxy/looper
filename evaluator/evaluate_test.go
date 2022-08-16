package evaluator

import (
	"testing"

	"github.com/Ranxy/looper/bind"
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
			vm := make(map[syntax.VariableSymbol]any)

			ev := NewEvaluater(boundTree.Statements, vm)
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
			vm := make(map[syntax.VariableSymbol]any)

			ev := NewEvaluater(boundTree.Statements, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEvaluate_variable(t *testing.T) {
	vm := make(map[syntax.VariableSymbol]any)

	var bt *bind.BoundGlobalScope

	bt = ev_variable(bt, vm, t, "var a = 1+1", int64(2))
	bt = ev_variable(bt, vm, t, "a=3", int64(3))
	bt = ev_variable(bt, vm, t, "-a", int64(-3))
	bt = ev_variable(bt, vm, t, "a+2", int64(5))

	bt = ev_variable(bt, vm, t, "a==3", true)

	bt = ev_variable(bt, vm, t, "a=a+a+1", int64(7))
	bt = ev_variable(bt, vm, t, "var b = false", false)
	bt = ev_variable(bt, vm, t, "b=(a==7)", true)
	_ = ev_variable(bt, vm, t, "b==false", false)

}

func ev_variable(previous *bind.BoundGlobalScope, vm map[syntax.VariableSymbol]any, t *testing.T, text string, want any) *bind.BoundGlobalScope {
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.ParseToTree(textSource)
	boundTree := bind.BindGlobalScope(previous, tree.Root)
	if len(boundTree.Diagnostic.List) != 0 {
		boundTree.Diagnostic.Print(text)
		t.FailNow()
	}
	ev := NewEvaluater(boundTree.Statements, vm)

	got := ev.Evaluate()
	require.Equal(t, want, got)

	return boundTree
}
