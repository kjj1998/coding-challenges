package count

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var BUFFER_SIZE int = 32 * 1024

func CountNumberOfBytes(file *os.File) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0

	for {
		c, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		count += c
	}

	return count
}

func CountNumberOfBytesFromStdin(reader *bytes.Reader) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0

	for {
		c, err := reader.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		count += c
	}

	return count
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

func CountNumberOfLinesFromStdin(reader *bytes.Reader) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0
	newLineSep := []byte{'\n'}

	for {
		c, err := reader.Read(buffer)

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

func CountNumberOfWordsFromStdin(reader *bytes.Reader) int {
	buffer := make([]byte, BUFFER_SIZE)
	count := 0
	inWord := false

	for {
		c, err := reader.Read(buffer)

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

func CountNumberOfCharacters(file *os.File) int {
	reader := bufio.NewReader(file)
	count := 0

	for {
		r, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if r != utf8.RuneError {
			count++
		}
	}

	return count
}

func CountNumberOfCharactersFromStdin(file *bytes.Reader) int {
	reader := bufio.NewReader(file)
	count := 0

	for {
		r, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		if r != utf8.RuneError {
			count++
		}
	}

	return count
}
