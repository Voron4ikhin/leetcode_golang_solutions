package _018_4_sum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{name: "Example 1", nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
