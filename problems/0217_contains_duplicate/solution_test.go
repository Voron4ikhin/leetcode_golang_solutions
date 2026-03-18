package _217_contains_duplicate

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContainsDuplicate result is = %v, want %v", got, tt.want)
			}
		})
	}
}
