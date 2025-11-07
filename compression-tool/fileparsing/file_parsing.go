package fileparsing

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func check(err error, filePath string) error {
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s: does not exist", filePath)
	}

	if errors.Is(err, os.ErrPermission) {
		return fmt.Errorf("%s: permission denied", filePath)
	}

	return fmt.Errorf("%s: %s", filePath, err.Error())
}

func ParseFile(filePath string) (map[rune]int, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, check(err, filePath)
	}

	text := string(data)
	words := strings.Fields(text)
	frequency := map[rune]int{}

	for _, word := range words {
		for _, ch := range word {
			_, ok := frequency[ch]

			if !ok {
				frequency[ch] = 1
			} else {
				frequency[ch]++
			}
		}
	}

	return frequency, nil
}
