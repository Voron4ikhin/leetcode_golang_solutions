package _049_group_anagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}, // может вернуться в любом порядке
		},
		{
			name: "Example 2",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "Example 3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupAnagram(tt.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagram result is = %v, want %v", got, tt.want)
			}
		})
	}
}
