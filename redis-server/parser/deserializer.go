package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Deserialize(bytes []byte) ([]string, error) {
	switch bytes[0] {
	case '+':
		result := ParseSimpleString(bytes)
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

func ParseSimpleString(bytes []byte) string {
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

func ReadSimpleString(reader *bufio.Reader) ([]byte, error) {
	simpleString, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	simpleString = bytes.TrimSpace(simpleString)

	return simpleString, nil
}

func ReadBulkString(reader *bufio.Reader) ([]byte, error) {
	lengthLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	length, err := strconv.Atoi(strings.TrimSpace(lengthLine))
	if err != nil {
		return nil, fmt.Errorf("non integer value given for bulk string length: %s", lengthLine)
	}

	bulkStringLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	bulkString := strings.TrimSpace(bulkStringLine)

	if length != len(bulkString) {
		return nil, fmt.Errorf("incorrect bulk string byte length given: %d", length)
	}

	return []byte(bulkString), nil
}

func ReadInteger(reader *bufio.Reader) ([]byte, error) {
	integerString, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	integerString = bytes.TrimSpace(integerString)

	if _, err := strconv.Atoi(string(integerString)); err != nil {
		return nil, fmt.Errorf("non integer value given: %s", integerString)
	}

	return integerString, nil
}

func DeserializeArray(reader *bufio.Reader) ([]any, error) {
	lengthLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	length, err := strconv.Atoi(strings.TrimSpace(lengthLine))
	if err != nil {
		return nil, fmt.Errorf("non integer value given for array length: %s", lengthLine)
	}

	elemCount := 0
	array := []any{}
	for {
		_, err := reader.Peek(1)
		if err == io.EOF {
			break
		}

		typeByte, err := reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("Unable to read byte: %s", err)
		}

		if typeByte == '*' {
			// Deserialze nested arrays recursively
			nestedArray, err := DeserializeArray(reader)
			if err != nil {
				return nil, err
			}
			array = append(array, nestedArray)
		} else {
			element, err := DeserializeSelector(typeByte, reader)
			if err != nil {
				return nil, err
			}
			array = append(array, element)
		}

		elemCount++
	}

	if length != len(array) {
		return nil, fmt.Errorf("incorrect number of elements given: %d", length)
	}

	return array, nil
}

func DeserializeSelector(typeByte byte, reader *bufio.Reader) ([]byte, error) {
	switch typeByte {
	case '+':
		// simple string
		simpleString, err := ReadSimpleString(reader)
		if err != nil {
			return nil, err
		}
		return simpleString, err
	case '$':
		// bulk string
		bulkString, err := ReadBulkString(reader)
		if err != nil {
			return nil, err
		}
		return bulkString, err
	case ':':
		// integer
		integer, err := ReadInteger(reader)
		if err != nil {
			return nil, err
		}
		return integer, err
	default:
		return nil, nil
	}
}
