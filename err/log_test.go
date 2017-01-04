package err

import (
	"testing"
)

func testFatalf(t *testing.T) {
	err := 1
	if err != 0 {
		err.Fatalf("testFatal: %v\n", err)
	}
}
