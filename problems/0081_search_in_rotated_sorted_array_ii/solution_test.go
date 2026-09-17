package _081_search_in_rotated_sorted_array_ii

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{name: "Example 1", nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0, want: true},
		{name: "Example 2", nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
