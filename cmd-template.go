// Command depend lists all installed packages that are dependant on input package[s]
package main

import (
	"flag"
	"fmt"
	"os"
)

var usage = `package [package ...]

List all installed packages that are dependant on input package[s]
`

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s %s\n", os.Args[0], usage)
	}
	flag.Parse()
}
