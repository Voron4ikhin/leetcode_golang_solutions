package _056_merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{name: "Example 1", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
