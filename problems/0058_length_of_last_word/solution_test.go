package _058_length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "Example 1", s: "Hello World", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
