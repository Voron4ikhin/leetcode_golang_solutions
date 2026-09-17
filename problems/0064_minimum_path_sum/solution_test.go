package _064_minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{name: "Example 1", grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum(tt.grid); got != tt.want {
				t.Errorf("MinPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
