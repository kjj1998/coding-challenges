package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	count "github.com/kjj1998/coding-challenges/wc/count"
)

var BUFFER_SIZE int = 32 * 1024

func main() {
	bytesFlag := flag.Bool("c", false, "Count the number of bytes in the file")
	linesFlag := flag.Bool("l", false, "Count the number of lines in the file")
	wordsFlag := flag.Bool("w", false, "Count the number of words in the file")

	flag.Parse()
	files := flag.Args()

	if len(files) != 1 {
		log.Fatal("No file supplied\n")
	}
	file, err := os.Open(files[0])

	if err != nil {
		panic(err)
	}

	if *bytesFlag {
		numberOfBytes := count.CountNumberOfBytes(file)
		fmt.Printf("%d %s\n", numberOfBytes, files[0])
	} else if *linesFlag {
		numberOfLines := count.CountNumberOfLines(file)
		fmt.Printf("%d %s\n", numberOfLines, files[0])
	} else if *wordsFlag {
		numberOfWords := count.CountNumberOfWords(file)
		fmt.Printf("%d %s\n", numberOfWords, files[0])
	}
}
