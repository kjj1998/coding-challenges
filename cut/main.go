package main

import (
	"flag"
	"fmt"

	cut "github.com/kjj1998/coding-challenges/cut/tool"
)

func main() {
	fieldPos := flag.String("f", "1", `The position(s) of the field to read, defaults to 1 for the first field. 
	To read multiple fields, input them separated by commas as in 1,2 or separated by whitespaces as in \"1 2 3\"`)
	delimiter := flag.String("d", "\t", "the delimiter used to separate the text in the file, defaults to tab")

	flag.Parse()

	// Get file name from positional arguments
	var file string
	if len(flag.Args()) > 0 {
		file = flag.Args()[0]
	}

	fieldPositions := cut.ParseFieldPos(fieldPos)
	cutValues := cut.Cut(fieldPositions, delimiter, &file)
	maxLen := 0

	for _, value := range cutValues {
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}

	for i := 0; i < maxLen; i++ {
		for j, value := range cutValues {
			if len(value) > i {
				fmt.Printf("%s", value[i])

				if j < len(cutValues)-1 {
					fmt.Printf("%s", *delimiter)
				}
			}
		}
		fmt.Println()
	}
}
