package _073_set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("SetZeroes() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
