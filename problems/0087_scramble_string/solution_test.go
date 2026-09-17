package _087_scramble_string

import "testing"

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{name: "Example 1", s1: "great", s2: "rgeat", want: true},
		{name: "Example 2", s1: "abcde", s2: "caebd", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsScramble(tt.s1, tt.s2); got != tt.want {
				t.Errorf("IsScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
