package _068_text_justification

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "Example 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want:     []string{"This    is    an", "example  of text", "justification.  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullJustify(tt.words, tt.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
