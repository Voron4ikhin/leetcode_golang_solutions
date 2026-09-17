package _063_unique_paths_ii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{name: "Example 1", obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
