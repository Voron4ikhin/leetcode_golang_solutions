package _008_string_to_integer_atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "Example 1", s: "42", want: 42},
		{name: "Example 2", s: "   -42", want: -42},
		{name: "Example 3", s: "4193 with words", want: 4193},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi(tt.s); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
