package _050_powx_n

import "testing"

func TestMyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{name: "Example 1", x: 2.0, n: 10, want: 1024.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyPow(tt.x, tt.n); got != tt.want {
				t.Errorf("MyPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
