package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	var r result

	start := time.Now()

	ticker := time.NewTicker(1 * time.Second).C

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r:
			return
		case <-ticker:
			log.Print("tick", r)
		}
	}

}

func first(ctx context.Context, urls []string) (*result, error) {

	results := make(chan result)

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}

func main() {

	urls := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	r, _ := first(context.Background(), urls)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("Quit anyway...", runtime.NumGoroutine(), "still running")

}
