package _089_gray_code

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{name: "Example 1", n: 2, want: []int{0, 1, 3, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrayCode(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
