package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
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

	var file *os.File = nil
	var reader *bytes.Reader = nil
	var fileName string
	var numberOfBytes int
	var numberOfLines int
	var numberOfWords int
	var numberOfCharacters int

	if len(files) == 1 {
		file, _ = os.Open(files[0])
		fileName = files[0]
	} else {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, os.Stdin)
		if err != nil {
			panic("stdin cannot be read")
		}

		reader = bytes.NewReader(buf.Bytes())
		fileName = ""
	}

	if *bytesFlag {
		if file != nil {
			numberOfBytes = count.CountNumberOfBytes(file)
		} else if reader != nil {
			numberOfBytes = count.CountNumberOfBytesFromStdin(reader)
		}
		fmt.Printf("%d %s\n", numberOfBytes, fileName)
	} else if *linesFlag {
		if file != nil {
			numberOfLines = count.CountNumberOfLines(file)
		} else if reader != nil {
			numberOfLines = count.CountNumberOfLinesFromStdin(reader)
		}
		fmt.Printf("%d %s\n", numberOfLines, fileName)
	} else if *wordsFlag {
		if file != nil {
			numberOfWords = count.CountNumberOfWords(file)
		} else if reader != nil {
			numberOfWords = count.CountNumberOfWordsFromStdin(reader)
		}
		fmt.Printf("%d %s\n", numberOfWords, fileName)
	} else if *charactersFlag {
		if file != nil {
			numberOfCharacters = count.CountNumberOfCharacters(file)
		} else if reader != nil {
			numberOfCharacters = count.CountNumberOfCharactersFromStdin(reader)
		}
		fmt.Printf("%d %s\n", numberOfCharacters, fileName)
	} else {
		if file != nil {
			numberOfBytes = count.CountNumberOfBytes(file)
			file.Seek(0, 0)
			numberOfLines = count.CountNumberOfLines(file)
			file.Seek(0, 0)
			numberOfWords = count.CountNumberOfWords(file)
		} else if reader != nil {
			numberOfBytes = count.CountNumberOfBytesFromStdin(reader)
			reader.Seek(0, io.SeekStart)
			numberOfLines = count.CountNumberOfLinesFromStdin(reader)
			reader.Seek(0, io.SeekStart)
			numberOfWords = count.CountNumberOfWordsFromStdin(reader)
		}

		fmt.Printf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, fileName)
	}
}
