package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func Deserialize(bytes []byte) ([]string, error) {
	switch bytes[0] {
	case '+':
		result := parseSimpleString(bytes)
		return []string{result}, nil
	case '-':
		result, err := parseError(bytes)
		if err != nil {
			return []string{}, err
		}
		return []string{result}, nil
	case ':':
		result, err := parseInteger(bytes)
		if err != nil {
			return []string{}, err
		}
		return []string{result}, nil
	case '$':
		result, _, err := parseBulkString(bytes)
		if err != nil {
			return []string{}, err
		}
		return []string{result}, nil
	case '*':
		result, err := parseArray(bytes)
		if err != nil {
			return []string{}, err
		}
		return result, nil
	}

	return []string{}, nil
}

func parseSimpleString(bytes []byte) string {
	endIndex := 0
	for i := range bytes {
		if bytes[i] != '\r' {
			endIndex++
		} else {
			break
		}
	}

	return string(bytes[1:endIndex])
}

func parseError(bytes []byte) (string, error) {
	clrfIndex := 0
	for i := range bytes {
		if bytes[i] != '\r' {
			clrfIndex++
		} else {
			break
		}
	}

	errMessage := string(bytes[1:clrfIndex])
	if !strings.HasPrefix(errMessage, "ERR") {
		return "", fmt.Errorf("Error message must begin with ERR: %s", errMessage)
	}

	return errMessage, nil
}

func parseInteger(bytes []byte) (string, error) {
	endIndex := 0
	for i := range bytes {
		if bytes[i] != '\r' {
			endIndex++
		} else {
			break
		}
	}

	integerInput := string(bytes[1:endIndex])
	_, err := strconv.Atoi(integerInput)
	if err != nil {
		return "", fmt.Errorf("non integer input: %s", integerInput)
	}

	return integerInput, nil
}

func parseBulkString(bytes []byte) (string, int, error) {
	firstCrlfIndex := 0

	for i := range bytes {
		if bytes[i] != '\r' {
			firstCrlfIndex++
		} else {
			break
		}
	}

	bulkStringByteLengthInput := string(bytes[1:firstCrlfIndex])
	bulkStringByteLength, err := strconv.Atoi(bulkStringByteLengthInput)
	if err != nil {
		return "", -1, fmt.Errorf("non integer value given for bulk string length: %s", bulkStringByteLengthInput)
	}
	if bulkStringByteLength == -1 {
		return "null", -1, nil
	}

	secondClrfIndex := firstCrlfIndex + 2 + bulkStringByteLength
	if secondClrfIndex >= len(bytes) || bytes[secondClrfIndex] != '\r' {
		return "", -1, fmt.Errorf("incorrect bulk string length input: %d", bulkStringByteLength)
	}
	bulkString := string(bytes[firstCrlfIndex+2 : secondClrfIndex])

	return bulkString, firstCrlfIndex - 1, nil
}

func parseArray(bytes []byte) ([]string, error) {
	endIndex := 0
	for i := range bytes {
		if bytes[i] != '\r' {
			endIndex++
		} else {
			break
		}
	}

	arrayLength, err := strconv.Atoi(string(bytes[1:endIndex]))
	if err != nil {
		panic("error reading array length")
	}
	if arrayLength == -1 {
		return []string{"null"}, nil
	}

	results := []string{}
	arrayIndex := endIndex + 2
	for range arrayLength {
		if arrayIndex >= len(bytes) {
			return []string{}, fmt.Errorf("array length is incorrect: %d", arrayLength)
		}

		bulkString, elementCount, err := parseBulkString(bytes[arrayIndex:])
		if err != nil {
			return []string{}, err
		}

		results = append(results, bulkString)
		arrayIndex += elementCount + 2 + len(bulkString) + 2 + 1
	}

	if arrayIndex < len(bytes) && bytes[arrayIndex] == '$' {
		return []string{}, fmt.Errorf("array length is incorrect: %d", arrayLength)
	}

	return results, nil
}
