package _057_insert_interval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{name: "Example 1", intervals: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}, want: [][]int{{1, 5}, {6, 9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.intervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
