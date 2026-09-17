package _016_3_sum_closest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "Example 1", nums: []int{-1, 2, 1, -4}, target: 1, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSumClosest(tt.nums, tt.target); got != tt.want {
				t.Errorf("ThreeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
