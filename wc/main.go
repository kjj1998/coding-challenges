package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputFile := flag.String("-c", "files/test.txt", "a txt file")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}

	stat, _ := file.Stat()

	fmt.Printf("%d %s\n", stat.Size(), *inputFile)
}
