package main

import (
	"flag"
	"fmt"

	cut "github.com/kjj1998/coding-challenges/cut/tool"
)

func main() {
	fieldPos := flag.String("f", "1", "The position of the field to read, defaults to 1 for the first field")
	delimiter := flag.String("d", "\t", "the delimiter used to separate the text in the file, defaults to tab")

	flag.Parse()

	// Get file name from positional arguments
	var file string
	if len(flag.Args()) > 0 {
		file = flag.Args()[0]
	}

	values := cut.Cut(fieldPos, delimiter, &file)

	for _, val := range values {
		fmt.Println(val)
	}
}
