package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// readInput читает одну строку из стандартного ввода
func readInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}

// splitLine разделяет строку на подстроки по пробелам
func splitLine(line string) []string {
	return strings.Fields(line)
}

// convertToInts преобразует подстроки в целые числа
func convertToInts(fields []string) ([]int, error) {
	var nums []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		if num < -10000 || num > 10000 {
			return nil, errors.New("the input number is greater than 10000 or less than -10000")
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// validateInput проверяет, что введено более одного числа
func validateInput(fields []string) error {
	if len(fields) <= 1 {
		return fmt.Errorf("you must enter more than one number")
	}
	return nil
}

// scanner читает строку, разделяет её на подстроки, преобразует их в целые числа и возвращает слайс целых чисел
func scanner() ([]int, error) {
	line, err := readInput()
	if err != nil {
		return nil, err
	}

	if line == "" {
		return nil, fmt.Errorf("Error: Empty input.")
	}

	fields := splitLine(line)

	if err := validateInput(fields); err != nil {
		return nil, err
	}

	nums, err := convertToInts(fields)
	if err != nil {
		return nil, err
	}

	return nums, nil
}

// arithmeticMean вычисляет среднее арифметическое слайса целых чисел
func arithmeticMean(nums []int) float64 {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return float64(sum) / float64(len(nums))
}

// median вычисляет медиану слайса целых чисел
func median(nums []int) float64 {
	length := len(nums)
	sort.Ints(nums) // Сортируем слайс

	if length%2 == 1 {
		// Если количество элементов нечетное, возвращаем средний элемент
		return float64(nums[length/2])
	} else {
		// Если количество элементов четное, возвращаем среднее значение двух средних элементов
		midIndex := length / 2
		return float64(nums[midIndex-1]+nums[midIndex]) / 2
	}
}

func mode(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	max_count := 0
	mode_num := nums[0]
	for num, count := range counts {
		if count > max_count || (count == max_count && num < mode_num) {
			max_count = count
			mode_num = num
		}
	}

	return mode_num
}

// standardDeviation вычисляет стандартное отклонение слайса целых чисел
func standardDeviation(nums []int, mean float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += (float64(num) - mean) * (float64(num) - mean)
	}
	return math.Sqrt(sum / float64(len(nums)))
}

// Пользователь может указать, какие метрики выводить, используя флаги командной строки.
// go run stat.go -mean=false -median=true -mode=true -sd=false
func main() {
	meanFlag := flag.Bool("mean", true, "Print mean")
	medianFlag := flag.Bool("median", true, "Print median")
	modeFlag := flag.Bool("mode", true, "Print mode")
	sdFlag := flag.Bool("sd", true, "Print standard deviation")

	flag.Parse()

	nums, err := scanner()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sort.Ints(nums)
	mean := arithmeticMean(nums)

	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", mean)
	}
	if *medianFlag {
		fmt.Printf("Median: %.2f\n", median(nums))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", mode(nums))
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", standardDeviation(nums, mean))
	}
}
