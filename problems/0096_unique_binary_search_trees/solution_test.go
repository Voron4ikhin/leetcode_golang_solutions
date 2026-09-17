package _096_unique_binary_search_trees

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "Example 1", n: 3, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees(tt.n); got != tt.want {
				t.Errorf("NumTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
