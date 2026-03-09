package _001_two_sum

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2",
			height: []int{1, 1},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxArea result is = %v, want %v", got, tt.want)
			}
		})
	}
}
