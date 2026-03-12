package _035_search_insert_position

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "Example 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "Example 3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInsert result is = %v, want %v", got, tt.want)
			}
		})
	}
}
