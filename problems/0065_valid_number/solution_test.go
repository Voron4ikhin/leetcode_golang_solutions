package _065_valid_number

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "Example 1", s: "0", want: true},
		{name: "Example 2", s: "e", want: false},
		{name: "Example 3", s: "2e10", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
