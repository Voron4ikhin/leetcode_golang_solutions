package _067_add_binary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "Example 1", a: "11", b: "1", want: "100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
