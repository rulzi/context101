package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "UserID", 123)

	performTask(ctx)
}

func performTask(ctx context.Context) {
	userID := ctx.Value("UserID")
	fmt.Println("Set User ID:", userID)
}
