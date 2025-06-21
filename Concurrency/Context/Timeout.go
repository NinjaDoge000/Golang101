package main

import (
	"context"
	"fmt"
	"time"
)

func workerTimeout(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped", ctx.Err())
			return

		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)

		}
	}
}

func Timeout() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		go workerTimeout(ctx)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
