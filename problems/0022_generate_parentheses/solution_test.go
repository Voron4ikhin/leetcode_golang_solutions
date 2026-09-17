package _022_generate_parentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "Example 1", n: 3, want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParenthesis(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
