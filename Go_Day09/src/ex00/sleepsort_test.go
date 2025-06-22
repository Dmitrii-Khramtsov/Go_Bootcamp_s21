package main

import (
	"testing"
)

func collect(ch <-chan int) []int {
	var result []int
	for val := range ch {
		result = append(result, val)
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSleepSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"emptySlice", []int{}, []int{}},
		{"singleElement", []int{42}, []int{42}},
		{"alreadySorted", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"reversed", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{"allEqual", []int{7, 7, 7, 7, 7}, []int{7, 7, 7, 7, 7}},
		{"negativesIncluded", []int{-5, 3, 0, -2, 7, -1}, []int{-5, -2, -1, 0, 3, 7}},
		{"withZeroes", []int{0, 0, 1, 2, 0}, []int{0, 0, 0, 1, 2}},
		{"randomSmall", []int{4, 1, 3, 2, 5}, []int{1, 2, 3, 4, 5}},
		{"randomLarge", []int{100, 3, 45, 2, 89, 1, 0}, []int{0, 1, 2, 3, 45, 89, 100}},
		{"maxAtStart", []int{100, 2, 3, 4, 5}, []int{2, 3, 4, 5, 100}},
		{"minAtEnd", []int{5, 4, 3, 2, 1, -10}, []int{-10, 1, 2, 3, 4, 5}},
		{"alternatingHighLow", []int{1, 100, 2, 99, 3, 98}, []int{1, 2, 3, 98, 99, 100}},
		{"duplicates", []int{5, 1, 2, 5, 3, 1}, []int{1, 1, 2, 3, 5, 5}},
		{"primeNumbers", []int{2, 3, 5, 7, 11, 13}, []int{2, 3, 5, 7, 11, 13}},
		{"powersOfTwo", []int{1, 2, 4, 8, 16, 32}, []int{1, 2, 4, 8, 16, 32}},
		{"fibonacciStart", []int{0, 1, 1, 2, 3, 5, 8, 13}, []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{"largeNumbers", []int{500, 1000, 999}, []int{500, 999, 1000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collect(sleepSort(tt.arr))
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
