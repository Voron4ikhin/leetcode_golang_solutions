package _028_find_the_index_of_the_first_occurrence_in_a_string

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{name: "Example 1", haystack: "sadbutsad", needle: "sad", want: 0},
		{name: "Example 2", haystack: "leetcode", needle: "leeto", want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
