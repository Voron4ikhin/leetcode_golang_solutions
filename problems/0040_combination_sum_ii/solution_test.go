package _040_combination_sum_ii

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "Example 1", candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum2(tt.candidates, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
