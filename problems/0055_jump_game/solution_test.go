package _055_jump_game

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "Example 1", nums: []int{2, 3, 1, 1, 4}, want: true},
		{name: "Example 2", nums: []int{3, 2, 1, 0, 4}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
