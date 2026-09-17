package _007_reverse_integer

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "Example 1",
			x:    123,
			want: 321,
		},
		{
			name: "Example 2",
			x:    -1230,
			want: -321,
		},
		{
			name: "Example 3",
			x:    2147483647,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse result is = %v, want %v", got, tt.want)
			}
		})
	}
}
