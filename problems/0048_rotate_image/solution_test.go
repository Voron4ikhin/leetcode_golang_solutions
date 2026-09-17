package _048_rotate_image

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("Rotate() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
