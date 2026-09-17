package _030_substring_with_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		{name: "Example 1", s: "barfoothefoobarman", words: []string{"foo", "bar"}, want: []int{0, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubstring(tt.s, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
