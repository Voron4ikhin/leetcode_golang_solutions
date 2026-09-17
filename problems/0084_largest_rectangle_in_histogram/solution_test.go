package _084_largest_rectangle_in_histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{name: "Example 1", heights: []int{2, 1, 5, 6, 2, 3}, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
