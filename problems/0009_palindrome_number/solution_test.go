package _009_palindrome_number

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "Example 1",
			x:    123,
			want: false,
		},
		{
			name: "Example 2",
			x:    1221,
			want: true,
		},
		{
			name: "Example 3",
			x:    431232134,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPalindrome result is = %v, want %v", got, tt.want)
			}
		})
	}
}
