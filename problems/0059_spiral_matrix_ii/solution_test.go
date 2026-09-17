package _059_spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{name: "Example 1", n: 3, want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
