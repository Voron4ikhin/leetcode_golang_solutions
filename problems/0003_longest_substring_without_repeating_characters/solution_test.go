package _003_longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example 3",
			s:    "pwwkew",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LengthOfLongestSubstring result is = %v, want %v", got, tt.want)
			}
		})
	}
}
