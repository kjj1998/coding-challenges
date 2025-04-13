package main

import (
	"flag"
	"fmt"
	"os"

	count "github.com/kjj1998/coding-challenges/wc/count"
)

var BUFFER_SIZE int = 32 * 1024

func main() {
	bytesFlag := flag.Bool("c", false, "Count the number of bytes in the file")
	linesFlag := flag.Bool("l", false, "Count the number of lines in the file")
	wordsFlag := flag.Bool("w", false, "Count the number of words in the file")
	charactersFlag := flag.Bool("m", false, "Count the number of characters in the file")

	flag.Parse()
	files := flag.Args()
	var reader *os.File
	var fileName string

	if len(files) == 1 {
		file, _ := os.Open(files[0])
		reader = file
		fileName = files[0]
	} else {
		reader = os.Stdin
		fileName = ""
	}

	if *bytesFlag {
		numberOfBytes := count.CountNumberOfBytes(reader)
		fmt.Printf("%d %s\n", numberOfBytes, fileName)
	} else if *linesFlag {
		numberOfLines := count.CountNumberOfLines(reader)
		fmt.Printf("%d %s\n", numberOfLines, fileName)
	} else if *wordsFlag {
		numberOfWords := count.CountNumberOfWords(reader)
		fmt.Printf("%d %s\n", numberOfWords, fileName)
	} else if *charactersFlag {
		numberOfCharacters := count.CountNumberOfCharacters(reader)
		fmt.Printf("%d %s\n", numberOfCharacters, fileName)
	} else {
		numberOfBytes := count.CountNumberOfBytes(reader)
		numberOfLines := count.CountNumberOfLines(reader)
		reader.Seek(0, 0)
		numberOfWords := count.CountNumberOfWords(reader)

		fmt.Printf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, fileName)
	}
}
