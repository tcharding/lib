package tricks

import (
	"log"
	"time"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = ^MaxInt
)

// Clear the noise on log output.
func fn() {
	log.Println("----------\n")
	defer log.Println("----------\n")

	// ...
}

// Create a timer, return function that stops the timer.
func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Since(t)
		log.Println(name, "took", d)
	}
}

// Example usage of timer.
func TimedFunc() {
	stop := StartTimer("TimedFunc")
	defer stop()

	// ...
}

//
// Parity
//

func IsEven(x int) bool {
	return x&-2 == x
}

func IsOdd(x int) bool {
	return !IsEven(x)
}
