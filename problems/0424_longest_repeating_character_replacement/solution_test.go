package _424_longest_repeating_character_replacement

import (
	"reflect"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "Example 2",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharacterReplacement(tt.s, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharacterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
