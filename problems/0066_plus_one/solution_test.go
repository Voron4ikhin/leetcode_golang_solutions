package _066_plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{name: "Example 1", digits: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{name: "Example 2", digits: []int{9, 9}, want: []int{1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
