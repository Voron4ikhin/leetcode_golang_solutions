package _027_remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{name: "Example 1", nums: []int{3, 2, 2, 3}, val: 3, want: 2},
		{name: "Example 2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.nums, tt.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
