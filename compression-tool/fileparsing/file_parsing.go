package fileparsing

import (
	"errors"
	"fmt"
	"os"
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

func ParseFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", check(err, filePath)
	}

	return string(data), nil
}
