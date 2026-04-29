package _849_maximize_distance_to_closest_person

import (
	"reflect"
	"testing"
)

func TestMaxDistToClosest(t *testing.T) {
	tests := []struct {
		name  string
		seats []int
		want  int
	}{
		{
			name:  "Example 1",
			seats: []int{1, 0, 0, 0, 1, 0, 1},
			want:  2,
		},
		{
			name:  "Example 2",
			seats: []int{1, 0, 0, 0},
			want:  3,
		},
		{
			name:  "Example 3",
			seats: []int{0, 1},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDistToClosest(tt.seats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
