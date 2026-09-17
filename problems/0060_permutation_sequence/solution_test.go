package _060_permutation_sequence

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{name: "Example 1", n: 3, k: 3, want: "213"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPermutation(tt.n, tt.k); got != tt.want {
				t.Errorf("GetPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
