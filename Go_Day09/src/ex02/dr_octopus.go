package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func multiplex(ctx context.Context, chs ...chan any) <-chan any {
	out := make(chan any, 100)
	wg := sync.WaitGroup{}

	output := func(ch chan any) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- val:
				}
			}
		}
	}

	for _, ch := range chs {
		wg.Add(1)
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	in1 := make(chan any, 32)
	in2 := make(chan any, 32)
	in3 := make(chan any, 32)

	go func() {
		defer close(in1)
		for i := range 5 {
			in1 <- i
		}
	}()

	go func() {
		defer close(in2)
		for i := range 5 {
			in2 <- 10 + i
		}
	}()

	go func() {
		defer close(in3)
		for i := range 5 {
			in3 <- 20 + i
		}
	}()

	out := multiplex(ctx, in1, in2, in3)

	for val := range out {
		fmt.Println(val)
	}
}
