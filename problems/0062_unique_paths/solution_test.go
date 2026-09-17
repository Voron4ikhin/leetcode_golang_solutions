package _062_unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{name: "Example 1", m: 3, n: 7, want: 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
