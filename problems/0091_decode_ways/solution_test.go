package _091_decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "Example 1", s: "12", want: 2},
		{name: "Example 2", s: "226", want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDecodings(tt.s); got != tt.want {
				t.Errorf("NumDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
