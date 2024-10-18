package main

import (
	"context"
	"log"
	"math/rand"
)

const keyID = "id"

func main() {
	// create a request ID as a random number
	requestID := rand.Intn(1000)

	// create a new context variable with a key value pair
	ctx := context.WithValue(context.Background(), keyID, requestID)
	operation1(ctx)
}

func operation1(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}
