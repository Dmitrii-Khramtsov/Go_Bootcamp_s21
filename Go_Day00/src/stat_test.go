package main

import (
	"testing"
)

// TestArithmeticMean тестирует функцию arithmeticMean.
func TestArithmeticMean(t *testing.T) {
	tests := []struct {
		nums   []int   // Входные данные
		expect float64 // Ожидаемый результат
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{10, 20, 30, 40, 50}, 30.0},
		{[]int{1, 1, 1, 1, 1}, 1.0},
	}

	for _, tt := range tests {
		result := arithmeticMean(tt.nums) // Вызов функции arithmeticMean
		if result != tt.expect {
			t.Errorf("arithmeticMean(%v) = %v; want %v", tt.nums, result, tt.expect) // Проверка результата
		}
	}
}

// TestMedian тестирует функцию median.
func TestMedian(t *testing.T) {
	tests := []struct {
		nums   []int   // Входные данные
		expect float64 // Ожидаемый результат
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{10, 20, 30, 40, 50}, 30.0},
		{[]int{1, 1, 1, 1, 1}, 1.0},
		{[]int{1, 2, 3, 4}, 2.5},
	}

	for _, tt := range tests {
		result := median(tt.nums) // Вызов функции median
		if result != tt.expect {
			t.Errorf("median(%v) = %v; want %v", tt.nums, result, tt.expect) // Проверка результата
		}
	}
}

// TestMode тестирует функцию mode.
func TestMode(t *testing.T) {
	tests := []struct {
		nums   []int // Входные данные
		expect int   // Ожидаемый результат
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 2, 3, 3, 4, 4, 4}, 4},
	}

	for _, tt := range tests {
		result := mode(tt.nums) // Вызов функции mode
		if result != tt.expect {
			t.Errorf("mode(%v) = %v; want %v", tt.nums, result, tt.expect) // Проверка результата
		}
	}
}

// TestStandardDeviation тестирует функцию standardDeviation.
func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		nums   []int   // Входные данные
		mean   float64 // Среднее значение
		expect float64 // Ожидаемый результат
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, 1.4142135623730951},
		{[]int{10, 20, 30, 40, 50}, 30.0, 14.142135623730951},
		{[]int{1, 1, 1, 1, 1}, 1.0, 0.0},
	}

	for _, tt := range tests {
		result := standardDeviation(tt.nums, tt.mean) // Вызов функции standardDeviation
		if result != tt.expect {
			t.Errorf("standardDeviation(%v, %v) = %v; want %v", tt.nums, tt.mean, result, tt.expect) // Проверка результата
		}
	}
}
