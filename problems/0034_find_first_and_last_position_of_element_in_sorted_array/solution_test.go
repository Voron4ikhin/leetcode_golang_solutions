package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "Example 1", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "Example 2", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
