package _013_roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "Example 1", s: "III", want: 3},
		{name: "Example 2", s: "LVIII", want: 58},
		{name: "Example 3", s: "MCMXCIV", want: 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
