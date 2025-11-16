package main

import (
	"flag"
	"fmt"

	"github.com/kjj1998/coding-challenges/cut/tool"
)

func main() {
	fieldPos := flag.String("f", "1", "The position of the field to read, defaults to 1 for the first field")
	delimiter := flag.String("d", "\t", "the delimiter used to separate the text in the file, defaults to tab")
	file := flag.String("file", "", "The name of the file to read from, defaults to empty string")

	flag.Parse()

	values := tool.Cut(fieldPos, delimiter, file)

	for _, val := range values {
		fmt.Println(val)
	}
}
