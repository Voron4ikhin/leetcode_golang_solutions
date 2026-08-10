package _005_longest_palindromic_substring

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "Example 2",
			s:    "cbbd",
			want: "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestPalindrome result is = %v, want %v", got, tt.want)
			}
		})
	}
}
