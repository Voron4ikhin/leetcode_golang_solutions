package _012_integer_to_roman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{name: "Example 1", num: 3, want: "III"},
		{name: "Example 2", num: 58, want: "LVIII"},
		{name: "Example 3", num: 1994, want: "MCMXCIV"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
