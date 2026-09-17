package _079_word_search

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.board, tt.word); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
