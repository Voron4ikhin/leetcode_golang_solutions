package _054_spiral_matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{name: "Example 1", matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
