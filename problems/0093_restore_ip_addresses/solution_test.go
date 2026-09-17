package _093_restore_ip_addresses

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{name: "Example 1", s: "25525511135", want: []string{"255.255.11.135", "255.255.111.35"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
