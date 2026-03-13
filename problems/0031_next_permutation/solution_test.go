package _031_next_permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextPermutation(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextPermutation result is = %v, want %v", got, tt.want)
			}
		})
	}
}
