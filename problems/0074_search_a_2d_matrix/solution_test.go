package _074_search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{name: "Example 1", matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3, want: true},
		{name: "Example 2", matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 13, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
