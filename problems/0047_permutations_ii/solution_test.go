package _047_permutations_ii

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "Example 1", nums: []int{1, 1, 2}, want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermuteUnique(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
