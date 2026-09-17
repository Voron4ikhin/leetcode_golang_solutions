package _046_permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "Example 1", nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
