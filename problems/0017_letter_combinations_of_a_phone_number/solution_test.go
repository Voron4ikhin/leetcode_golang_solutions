package _017_letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{name: "Example 1", digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "Example 2", digits: "", want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
