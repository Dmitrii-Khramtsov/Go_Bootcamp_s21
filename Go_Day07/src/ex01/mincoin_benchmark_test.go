package mincoins

import (
	"testing"
)

func BenchmarkMinCoins(b *testing.B) {
	coins := []int{1, 3, 4, 7, 13, 15, 21, 37, 52, 78, 91, 101}
val := 100000

	for i := 0; i < b.N; i++ {
		minCoins(val, coins)
	}
}

func BenchmarkMinCoinsOptimized(b *testing.B) {
	coins := []int{1, 3, 4, 7, 13, 15, 21, 37, 52, 78, 91, 101}
val := 100000

	for i := 0; i < b.N; i++ {
		minCoinsOptimized(val, coins)
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	coins := []int{1, 3, 4, 7, 13, 15, 21, 37, 52, 78, 91, 101}
val := 100000

	for i := 0; i < b.N; i++ {
		minCoins2(val, coins)
	}
}

func BenchmarkMinCoins2Optimized(b *testing.B) {
	coins := []int{1, 3, 4, 7, 13, 15, 21, 37, 52, 78, 91, 101}
val := 100000

	for i := 0; i < b.N; i++ {
		minCoins2Optimized(val, coins)
	}
}

// go test -bench=. -cpuprofile=cpu.prof -benchmem

// go tool pprof -top cpu.prof
// go tool pprof -top cpu.prof > top10.txt
