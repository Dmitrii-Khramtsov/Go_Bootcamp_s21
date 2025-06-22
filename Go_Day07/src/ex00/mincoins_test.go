package mincoins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCoins_GreedyFails(t *testing.T) {
	coins := []int{1, 3, 4}
	val := 6

	greedy := minCoins(val, coins)
	correct := []int{3, 3}
	
	require.NotEqual(t, correct, greedy, "Greedy algorithm fails here")
}

func TestMinCoins2_Correctness(t *testing.T) {
	tests := []struct {
		coins []int
		val   int
		want  []int
	}{
		{[]int{1, 3, 4}, 6, []int{3, 3}},
		{[]int{1, 5, 10, 50, 100, 500, 1000}, 58, []int{50, 5, 1, 1, 1}},
		{[]int{1, 3, 4, 7, 13, 15}, 14, []int{7, 7}},
		{[]int{}, 10, []int{}},
		{[]int{1}, 3, []int{1, 1, 1}},
	}

	for _, tt := range tests {
		got := minCoins2Optimized(tt.val, tt.coins)

		if len(tt.coins) == 0 {
			require.Empty(t, got, "empty coins should return empty result")
			continue
		}

		sum := 0
		for _, coin := range got {
			sum += coin
		}

		require.Equal(t, tt.val, sum, "sum mismatch")
		require.Equal(t, len(got), len(tt.want), "amount mismatch")
		require.NotNil(t, got)
	}
}
