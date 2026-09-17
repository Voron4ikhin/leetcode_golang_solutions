package _052_n_queens_ii

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "Example 1", n: 4, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalNQueens(tt.n); got != tt.want {
				t.Errorf("TotalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
