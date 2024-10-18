package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	status := performTask(ctx)

	fmt.Println(status)
}

func performTask(ctx context.Context) bool {
	done := make(chan bool)

	go func() {
		// Run Task
		time.Sleep(2 * time.Second)
		fmt.Println("Run Task")

		done <- true
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Perform Task timed out")
		return false
	case <-done:
		return true
	}
}
