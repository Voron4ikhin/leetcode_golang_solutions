package _077_combinations

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{name: "Example 1", n: 4, k: 2, want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combine(tt.n, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
