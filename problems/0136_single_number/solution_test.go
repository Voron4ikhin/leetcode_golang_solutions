package _136_single_number

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "Example 3",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleNumber result is = %v, want %v", got, tt.want)
			}
		})
	}
}
