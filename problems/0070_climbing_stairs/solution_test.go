package _070_climbing_stairs

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    2,
			want: 2,
		},
		{
			name: "Example 2",
			n:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClimbStairs result is = %v, want %v", got, tt.want)
			}
		})
	}
}
