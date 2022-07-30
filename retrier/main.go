package main

import (
	"fmt"
	"time"
)

// RetryGenerics retries a function until it returns success or the retry times is reached.
// The retry delay is in milliseconds.
// The fn MUST accept V as parameter and return K and error
func RetryGenerics[V, K any](fn func(e V) (K, error), params V, retryTimes int, retryDelayMs int) (K, bool) {
	var r K
	for i := 0; i < retryTimes; i++ {
		r, err := fn(params)
		if err == nil {
			return r, true
		}
		fmt.Printf("Error calling fn: %+v\r\n", err)
		time.Sleep(time.Duration(retryDelayMs) * time.Millisecond)
	}
	return r, false
}

var timesCalled = 0

// PrintNameGenericWithError prints the name and returns an error if the timesCalled is less than 5.
// this is part of the retry example.
func PrintNameGenericWithError(name string) (string, error) {
	timesCalled++
	if timesCalled >= 5 {
		fmt.Println(name)
		return name, nil
	}

	return "", fmt.Errorf("not yet")
}

func main() {
	res, success := RetryGenerics(PrintNameGenericWithError, "edu generic", 7, 1000)
	if success {
		fmt.Println("Success: ", res)
	} else {
		fmt.Println("Failed")
	}
}
