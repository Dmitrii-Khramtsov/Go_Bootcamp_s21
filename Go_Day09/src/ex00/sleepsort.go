package main

import (
	"fmt"
	"sync"
	"time"
)

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func sleepSort(arr []int) <-chan int {
	if len(arr) == 0 {
		ch := make(chan int)
		close(ch)
		return ch
	}

	minVal := findMin(arr)
	ch := make(chan int, len(arr))
	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			delay := time.Duration(v-minVal) * 20 * time.Millisecond
			time.Sleep(delay)
			ch <- v
		}(val)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	arr := []int{-5, 3, 0, -2, 700, -100}
	ch := sleepSort(arr)
	for val := range ch {
		fmt.Println(val)
	}
}
