package utils

import (
	"fmt"
	"os"
)

// fatalErr: print format string to standard out including err and exit
func FatalErr(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", err)
	os.Exit(1)
}

// fatalf: print format string to standard out and exit
func Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}
