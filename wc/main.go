package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	bytesFlag := flag.Bool("c", false, "Count the number of bytes in the file")
	flag.Parse()

	files := flag.Args()

	file, err := os.Open(files[0])
	if err != nil {
		panic(err)
	}

	stat, _ := file.Stat()

	if *bytesFlag {
		fmt.Printf("%d %s\n", stat.Size(), files[0])
	}
}
