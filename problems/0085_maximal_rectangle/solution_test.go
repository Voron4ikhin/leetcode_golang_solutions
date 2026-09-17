package _085_maximal_rectangle

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "Example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximalRectangle(tt.matrix); got != tt.want {
				t.Errorf("MaximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
