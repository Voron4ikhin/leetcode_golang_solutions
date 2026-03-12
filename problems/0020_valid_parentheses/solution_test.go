package _020_valid_parentheses

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "Example 4",
			s:    "([])",
			want: true,
		},
		{
			name: "Example 5",
			s:    "([)]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsValid result is = %v, want %v", got, tt.want)
			}
		})
	}
}
