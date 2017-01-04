package tricks

import (
	"log"
	"testing"
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

// Skip long tests, `go test -v -short`.
func Test(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	// rest of tests
}

// IsEven: true if x is even.
func IsEven(x int) bool {
	return x&-2 == x
}

// IsOdd: true if x is odd.
func IsOdd(x int) bool {
	return !IsEven(x)
}

// Mod: Integer modulo.
func Mod(a, b int) int {
	return a - (a/b)*b
}

// Abs: absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}
