package _071_simplify_path

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{name: "Example 1", path: "/home/", want: "/home"},
		{name: "Example 2", path: "/../", want: "/"},
		{name: "Example 3", path: "/home//foo/", want: "/home/foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyPath(tt.path); got != tt.want {
				t.Errorf("SimplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
