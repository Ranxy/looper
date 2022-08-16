package texts

import (
	"testing"
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
			got := NewTextSource(tt.args.text)
			if len(got.Lines) != tt.want {
				t.Errorf("TextSource Len = %v, want %v", len(got.Lines), tt.want)
			}
		})
	}
}
