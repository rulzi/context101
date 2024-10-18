package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	perform, err := performCalc(ctx, 10, 20)

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	log.Println("perform success. result:", *perform)
}

func performCalc(ctx context.Context, number1 int32, number2 int32) (*int32, error) {

	type Result struct {
		calculate *int32
		err       error
	}

	result := make(chan Result)

	// call Thirparty in goroutine
	go func() {
		calculate, err := thirdPartyPerformCalc(number1, number2)

		res := Result{
			calculate: calculate,
			err:       err,
		}

		result <- res

	}()

	select {
	case <-ctx.Done():
		fmt.Println("Perform Task timed out")
		return nil, errors.New("Perform Task timed out")
	case result := <-result:
		return result.calculate, result.err
	}

}

func thirdPartyPerformCalc(number1 int32, number2 int32) (*int32, error) {
	// time.Sleep(1 * time.Second)
	// return nil, errors.New("Fail")

	operation := number1 + number2
	return &operation, nil
}
