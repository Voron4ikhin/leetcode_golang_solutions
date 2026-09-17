package _076_minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{name: "Example 1", s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
