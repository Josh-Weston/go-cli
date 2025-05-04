package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// An alternative approach to declaring the boolean value
	// var useLines bool
	// flag.BoolVar(&useLines, "l", false, "Count lines")

	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes)) // Dereferences the lines pointer
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords) // Split using the bufio.ScanWords function
	}

	wc := 0
	for scanner.Scan() {
		if countBytes {
			wc += len(scanner.Bytes())
		} else {
			wc++
		}
	}
	return wc
}
