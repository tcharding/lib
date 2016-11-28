package tricks

import (
	"fmt"
	"testing"
)

const verbose = false

func TestTimer(t *testing.T) {
	if verbose {
		stop := StartTimer("TimedFunc")
		defer func() {
			stop()
			fmt.Println()
		}()

		big := 1 << 28
		for i := 0; i < big; i++ {
			// busy loop
			a := i
			b := a
			a = b
		}
	}
}

func TestParity(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{0, true},
		{2, true},
		{64, true},

		{1, false},
		{3, false},
		{65, false},
		{MaxInt, false},
	}
	// Test IsEven.
	for _, test := range tests {
		if got := IsEven(test.input); got != test.want {
			t.Errorf("IsEven(%d) = %v want: %v\n", test.input, got, test.want)
		}
	}
	// Test IsOdd.
	for _, test := range tests {
		if got := IsOdd(test.input); got == test.want {
			t.Errorf("IsOdd(%d) = %v want: %v\n", test.input, got, !test.want)
		}
	}
}
