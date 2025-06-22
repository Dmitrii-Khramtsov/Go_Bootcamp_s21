package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMultiplexWithSlowAndUnclosedChannels(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	in1 := make(chan any)
	in2 := make(chan any)
	in3 := make(chan any) // никогда не закроется

	go func() {
			defer close(in1)
			for i := range 5 {
					in1 <- i
			}
	}()

	go func() {
			defer close(in2)
			for i := range 5 {
					time.Sleep(250 * time.Millisecond) // задержка 250ms
					in2 <- 100 + i
			}
	}()

	go func() {
			for i := range 5 {
					in3 <- 200 + i
			}
	}()

	out := multiplex(ctx, in1, in2, in3)

	received := make(map[any]bool)
	for val := range out {
			received[val] = true
			switch v := val.(int); {
			case v < 100:
				fmt.Printf("\033[34m[CH1]\033[0m got: %v\n", v)
			case v < 200:
				fmt.Printf("\033[32m[CH2]\033[0m got: %v\n", v)
			default:
				fmt.Printf("\033[35m[CH3]\033[0m got: %v\n", v)
			}
	}

	for i := range 5 {
			if !received[i] {
					t.Errorf("missing value from CH1: %v", i)
			}
	}

	receivedFromSlow := 0
	for i := 100; i < 105; i++ {
			if received[i] {
					receivedFromSlow++
			}
	}
	if receivedFromSlow < 3 { // должны успеть минимум 3 значения (1000ms / 250ms = 4, но учитываем погрешности)
			t.Errorf("expected at least 3 values from CH2, got %d", receivedFromSlow)
	}

	for i := 200; i < 205; i++ {
			if !received[i] {
					t.Errorf("missing value from CH3: %v", i)
			}
	}
}
