package _078_subsets

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "Example 1", nums: []int{1, 2, 3}, want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
