package main

import (
	"context"
	"fmt"
	"time"
)

func workerDeadline(ctx context.Context) {

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

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	for i := 0; i < 3; i++ {
		go workerDeadline(ctx)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
