package _033_search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "Example 1", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{name: "Example 2", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
