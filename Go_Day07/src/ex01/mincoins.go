package mincoins

import "sort"

func coinsNormalize(coins []int) []int {
	unique := make(map[int]struct{}, len(coins))

	for _, coin := range coins {
		if _, ok := unique[coin]; !ok {
			unique[coin] = struct{}{}
		}
	}

	res := make([]int, 0, len(unique))
	for val := range unique {
		res = append(res, val)
	}

	// sort.Sort(sort.Reverse(sort.IntSlice(res)))

	return res
}

// Greedy
func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoinsOptimized(val int, coins []int) []int {
	coinsNorm := coinsNormalize(coins)
	sort.Sort(sort.Reverse(sort.IntSlice(coinsNorm)))

	dp := make([]int, 0, len(coinsNorm))

	tmp := val
	for _, coin := range coinsNorm {
		if coin <= tmp {
			count := int(tmp / coin)
			tmp -= coin * count
			for range count {
        dp = append(dp, coin)
    }
		}
	}

	return dp
}

// Memo
func minCoins2(val int, coins []int) []int {
	if val == 0 {
		return []int{}
	}
	if val < 0 {
		return nil
	}

	coinsNorm := coinsNormalize(coins)
	memo := make(map[int][]int, len(coins)) // ключ - остаток, разложение с мин. кол. coins

	var recurseRemainder func(int) []int
	recurseRemainder = func(remaind int) []int {
		if remaind == 0 {
			return []int{}
		}
		if remaind < 0 {
			return nil
		}
		if v, ok := memo[remaind]; ok {
			return v
		}

		var best []int

		for _, coin := range coinsNorm {
			if coin <= remaind {
				subResult := recurseRemainder(remaind - coin)
				if subResult == nil { // nil - флаг, невозможно собрать сумму для этого остатка (remaind < 0) - продолжаем новую итерацию
					continue
				}
				candidate := append([]int{coin}, subResult...)

				if best == nil || len(candidate) < len(best) {
					best = candidate
				}
			}
		}

		memo[remaind] = best
		return best
	}

	return recurseRemainder(val)
}

// DP
func minCoins2Optimized(val int, coins []int) []int {
	if val == 0 {
		return []int{}
	}
	if val < 0 {
		return nil
	}

	coinsNorm := coinsNormalize(coins)

	dp := make([][]int, val+1)
	dp[0] = []int{}

	for i := 1; i <= val; i++ {
		var best []int
		for _, coin := range coinsNorm {
			if coin <= i && dp[i-coin] != nil { // без инициализации dp[0] = []int{}, все dp[i] == nil
				candidate := append([]int{coin}, dp[i-coin]...)
				if best == nil || len(candidate) < len(best) {
					best = candidate
				}
			}
		}
		dp[i] = best
	}

	return dp[val]
}
