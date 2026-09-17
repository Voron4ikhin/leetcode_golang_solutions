package _044_wildcard_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{name: "Example 1", s: "aa", p: "a", want: false},
		{name: "Example 2", s: "aa", p: "*", want: true},
		{name: "Example 3", s: "cb", p: "?a", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
