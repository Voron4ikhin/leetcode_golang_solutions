package _014_longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{name: "Example 1", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "Example 2", strs: []string{"dog", "racecar", "car"}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
