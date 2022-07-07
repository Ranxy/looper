package evaluator

import (
	"fmt"
	"testing"

	"github.com/Ranxy/looper/bind"
	"github.com/Ranxy/looper/syntax"
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

			tree := syntax.NewParser(tt.text).Parse()

			bound := bind.NewBinder()
			boundTree := bound.BindExpression(tree.Root)
			if len(bound.Errors) != 0 {
				fmt.Println(bound.Errors)
				t.FailNow()
			}
			got := Evaluate(boundTree)
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

			tree := syntax.NewParser(tt.text).Parse()
			bound := bind.NewBinder()
			boundTree := bound.BindExpression(tree.Root)
			require.Zero(t, len(bound.Errors))

			got := Evaluate(boundTree)
			require.Equal(t, tt.want, got)
		})
	}
}
