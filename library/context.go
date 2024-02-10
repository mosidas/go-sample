package library

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Context() {
	fmt.Println("=== context ===")
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, &wg)
	time.Sleep(1500 * time.Millisecond)
	cancel()
	wg.Wait()
}

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("work")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
