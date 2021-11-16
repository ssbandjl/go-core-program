package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			log.Printf("hello")
		case <-ctx.Done():
			log.Printf("exit reson:%s", ctx.Err())
			return ctx.Err()
		}

	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
