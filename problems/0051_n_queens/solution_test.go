package _051_n_queens

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "Example 1",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveNQueens(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
