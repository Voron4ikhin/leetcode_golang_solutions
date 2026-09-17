package _038_count_and_say

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "Example 1", n: 1, want: "1"},
		{name: "Example 2", n: 4, want: "1211"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountAndSay(tt.n); got != tt.want {
				t.Errorf("CountAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
