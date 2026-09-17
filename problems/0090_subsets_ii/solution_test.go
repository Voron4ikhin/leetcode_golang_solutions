package _090_subsets_ii

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "Example 1", nums: []int{1, 2, 2}, want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetsWithDup(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
