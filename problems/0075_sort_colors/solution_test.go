package _075_sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Example 2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortColors(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortColors result is = %v, want %v", got, tt.want)
			}
		})
	}
}
