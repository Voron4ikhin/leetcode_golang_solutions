package _072_edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{name: "Example 1", word1: "horse", word2: "ros", want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
