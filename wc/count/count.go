package count

import (
	"bytes"
	"io"
	"os"
	"unicode"
)

var BUFFER_SIZE int = 32 * 1024

func CountNumberOfBytes(file *os.File) int {
	stat, _ := file.Stat()

	return int(stat.Size())
}

func CountNumberOfLines(file *os.File) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0
	newLineSep := []byte{'\n'}

	for {
		c, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		count += bytes.Count(buffer[:c], newLineSep)
	}

	return count
}

func CountNumberOfWords(file *os.File) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0
	inWord := false

	for {
		c, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		for _, b := range buffer[:c] {
			if unicode.IsSpace(rune(b)) {
				inWord = false
			} else if !inWord {
				inWord = true
				count++
			}
		}
	}

	return count
}
