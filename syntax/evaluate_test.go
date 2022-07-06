package syntax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		text    string
		want    int64
		wantErr bool
	}{
		{
			text:    "1 + 2",
			want:    3,
			wantErr: false,
		},
		{
			text:    "10 - 2 * 3",
			want:    4,
			wantErr: false,
		},
		{
			text:    "1 + 2 * (3 -1 )",
			want:    5,
			wantErr: false,
		},
		{
			text:    "-2",
			want:    -2,
			wantErr: false,
		},
		{
			text:    "1+-2",
			want:    -1,
			wantErr: false,
		},
		{
			text:    "-(1 * 2) -3",
			want:    -5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {

			tree := NewParser(tt.text).Parse()
			got, err := Evaluate(tree.Root)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
