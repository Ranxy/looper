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
			vm := bind.NewVariableManage()
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.NewParser(textSource).Parse()

			bound := bind.NewBinder(vm)
			boundTree := bound.BindExpression(tree.Root)
			if len(bound.Diagnostics.List) != 0 {
				bound.Diagnostics.Print(tt.text)
				t.FailNow()
			}
			ev := NewEvaluater(boundTree, vm)
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
			vm := bind.NewVariableManage()
			textSource := texts.NewTextSource([]rune(tt.text))
			tree := syntax.NewParser(textSource).Parse()
			bound := bind.NewBinder(vm)
			boundTree := bound.BindExpression(tree.Root)
			require.Zero(t, len(bound.Diagnostics.List))

			ev := NewEvaluater(boundTree, vm)
			got := ev.Evaluate()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEvaluate_variable(t *testing.T) {
	vm := bind.NewVariableManage()

	ev_variable(vm, t, "a=3", int64(3))
	ev_variable(vm, t, "-a", int64(-3))
	ev_variable(vm, t, "a+2", int64(5))

	ev_variable(vm, t, "a==3", true)

	ev_variable(vm, t, "a=a+a+1", int64(7))
	ev_variable(vm, t, "b=(a==7)", true)
	ev_variable(vm, t, "b==false", false)
}

func ev_variable(vm bind.VariableManage, t *testing.T, text string, want any) {
	textSource := texts.NewTextSource([]rune(text))
	tree := syntax.NewParser(textSource).Parse()
	bound := bind.NewBinder(vm)
	boundTree := bound.BindExpression(tree.Root)
	bound.Diagnostics.Print(text)
	require.Zero(t, len(bound.Diagnostics.List))

	ev := NewEvaluater(boundTree, vm)
	got := ev.Evaluate()
	require.Equal(t, want, got)
}
