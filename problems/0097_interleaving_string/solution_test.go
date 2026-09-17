package _097_interleaving_string

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{name: "Example 1", s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac", want: true},
		{name: "Example 2", s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("IsInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
