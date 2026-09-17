package _045_jump_game_ii

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "Example 1", nums: []int{2, 3, 1, 1, 4}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.nums); got != tt.want {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
