package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped")
			return

		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)

		}
	}
}

func cancel() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 3; i++ {
		go worker(ctx)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
