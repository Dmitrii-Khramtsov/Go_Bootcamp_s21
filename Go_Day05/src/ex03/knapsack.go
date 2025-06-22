package main

import "fmt"

type Present struct {
	Value int
	Size  int
}

func grabPresents(presents []Present, capacity int) []Present {
	dp := make([]int, capacity+1)
	selected := make([][]bool, len(presents)+1)

	for i := range selected {
		selected[i] = make([]bool, capacity+1)
	}

	for i, p := range presents {
		for w := capacity; w >= p.Size; w-- {
			if dp[w-p.Size]+p.Value > dp[w] {
				dp[w] = dp[w-p.Size] + p.Value
				selected[i+1][w] = true
			}
		}
		// fmt.Println(dp)
	}

	res := []Present{}
	w := capacity
	for i := len(presents) - 1; i >= 0; i-- {
		if selected[i+1][w] {
			res = append(res, presents[i])
			w -= presents[i].Size
		}
	}
	return res
}

func main() {
	presents_test := make([]Present, 4)
	presents_test[0].Size = 1
	presents_test[0].Value = 1500
	presents_test[1].Size = 4
	presents_test[1].Value = 3000
	presents_test[2].Size = 3
	presents_test[2].Value = 2000
	presents_test[3].Size = 1
	presents_test[3].Value = 2000

	fmt.Println(grabPresents(presents_test, 4))
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func grabPresents(presents []Present, capacity int) []Present {
// 	n := len(presents)

// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, capacity+1)
// 	}

// 	for i := 1; i <= n; i++ {
// 		for w := 0; w <= capacity; w++ {
// 			current := presents[i-1]
// 			if current.Size > w {
// 				dp[i][w] = dp[i-1][w]
// 			} else {
// 				dp[i][w] = max(dp[i-1][w], dp[i-1][w-current.Size]+current.Value)
// 			}
// 		}
// 	}

// 	fmt.Println(dp)

// 	maxTotalValue := []Present{}
// 	w := capacity
// 	for i := n; i > 0 && w > 0; i-- {
// 		if dp[i][w] != dp[i-1][w] {
// 			maxTotalValue = append(maxTotalValue, presents[i-1])
// 			w -= presents[i-1].Size
// 		}
// 	}

// 	return maxTotalValue
// }
