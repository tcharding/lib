// err implements simple error functionality
package err

import (
	"fmt"
	"os"
)

// fatalf is equivalent to fmt.Fprintf(os.Stderr, "...") followed by a call to os.Exit(1).
func fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

// Msg is equivalent to fmt.Fprintf(os.Stderr, "...").
func Msg(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
}
