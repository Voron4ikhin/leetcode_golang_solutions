package _069_sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "Example 1", x: 4, want: 2},
		{name: "Example 2", x: 8, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySqrt(tt.x); got != tt.want {
				t.Errorf("MySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
