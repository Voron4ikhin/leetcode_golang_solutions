package _032_longest_valid_parentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "Example 1", s: "(()", want: 2},
		{name: "Example 2", s: ")()())", want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestValidParentheses(tt.s); got != tt.want {
				t.Errorf("LongestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
