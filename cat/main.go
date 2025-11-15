package main

import (
	"flag"
	"fmt"
)

func main() {
	fieldPos := flag.String("f", "1", "The position of the field to read, defaults to 1 for the first field")
	file := flag.String("file", "", "The name of the file to read from, defaults to empty string")

	flag.Parse()

	fmt.Printf("fieldPos: %s, file: %s\n", *fieldPos, *file)
}
