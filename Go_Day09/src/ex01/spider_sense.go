package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const limit = 8

func fetchHTML(ctx context.Context, url string) (string, error) {
	ctx, close := context.WithTimeout(ctx, 5*time.Second)
	defer close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 200)
	n, _ := resp.Body.Read(buf)

	return string(buf[:n]), nil
}

func worker(ctx context.Context, in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
		case url, ok := <-in:
			if !ok {
				return
			}
			html, err := fetchHTML(ctx, url)
			if err != nil {
				fmt.Println("fetch error:", err)
				continue
			}
			select {
			case out <- html:
			case <-ctx.Done():
				return
			}
		} 
		return
	}
}

func startWorkers(in <-chan string, limiter int, ctx context.Context) <-chan string {
	out := make(chan string, 32)
	wg := sync.WaitGroup{}

	for range limiter {
		wg.Add(1)
		go worker(ctx, in, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func crawlWeb(in <-chan string, limiter int, ctx context.Context) <-chan string {
	return startWorkers(in, limiter, ctx)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	job := make(chan string, 32)

	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://httpbin.org/uuid",
		"https://httpbin.org/ip",
		"https://httpbin.org/user-agent",
		"https://httpbin.org/headers",
		"https://httpbin.org/delay/1",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/204",
	}

	go func() {
		defer close(job)
		for _, url := range urls {
			select {
			case <- ctx.Done():
				return
			case job <- url:
			}
		}
	}()

	for html := range crawlWeb(job, limit, ctx) {
		fmt.Println("--- Response ---")
		fmt.Println(html)
	}
}
