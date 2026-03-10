package _121_best_time_to_buy_sell_stocks

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Example 2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
