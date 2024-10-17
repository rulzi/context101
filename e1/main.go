package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go performTask(ctx)
	go performTask2(ctx)
	go performTask3(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}
}

func performTask(ctx context.Context) {
	ctxPerform, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	go performTask11(ctxPerform)
	go performTask12(ctxPerform)
	go performTask13(ctxPerform)

	select {
	case <-ctx.Done():
		fmt.Println("Perform Task 1 timed out")
		return
	}
}

func performTask2(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("Task completed successfully 2")
		}
	}
}

func performTask3(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("Task completed successfully 3")
		}
	}
}

func performTask11(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Perform Task 11 timed out")
			return
		case <-ticker.C:
			fmt.Println("Task completed successfully 11")
		}
	}
}

func performTask12(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Perform Task 12 timed out")
			return
		case <-ticker.C:
			fmt.Println("Task completed successfully 12")
		}
	}
}

func performTask13(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Perform Task 13 timed out")
			return
		case <-ticker.C:
			fmt.Println("Task completed successfully 13")
		}
	}
}
