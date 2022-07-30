package main

import (
	"fmt"
	"time"
)

// Retrier retries a function until it returns success or the retry times is reached.
// The retry delay is in milliseconds.
// The fn MUST not accept any parameter but it must return an error
func Retrier(fn func() error, retryTimes int, retryDelayMs int) bool {
	for i := 0; i < retryTimes; i++ {
		err := fn()
		if err == nil {
			return true
		}
		fmt.Printf("Error calling fn: %+v\r\n", err)
		time.Sleep(time.Duration(retryDelayMs) * time.Millisecond)
	}
	return false
}

var timesCalled = 0

// PrintNameWithError prints the name and returns an error if the timesCalled is less than 5.
// this is part of the retry example.
func PrintNameWithError(name string) error {
	timesCalled++
	if timesCalled >= 5 {
		fmt.Println(name)
		return nil
	}

	return fmt.Errorf("not yet")
}

func main() {
	fn := func() error {
		return PrintNameWithError("edu")
	}

	res := Retrier(fn, 7, 1000)
	if res {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
