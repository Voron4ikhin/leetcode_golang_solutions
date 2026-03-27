package _242_valid_anagram

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Example 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsAnagram result is = %v, want %v", got, tt.want)
			}
		})
	}
}
