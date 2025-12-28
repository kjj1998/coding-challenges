package parser

import (
	"bytes"
	"strconv"
)

func Deserialize(buf bytes.Buffer) ([]string, error) {
	bytes := buf.Bytes()

	switch bytes[0] {
	case '+':
		result := parseSimpleString(bytes)
		return []string{result}, nil
	case '-':
		result := parseError(bytes)
		return []string{result}, nil
	case ':':
		result := parseInteger(bytes)
		return []string{result}, nil
	case '$':
		result, _ := parseBulkString(bytes)
		return []string{result}, nil
	case '*':
		result := parseArray(bytes)
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

func parseError(bytes []byte) string {
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

func parseInteger(bytes []byte) string {
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

func parseBulkString(bytes []byte) (string, int) {
	endIndex := 0

	for i := range bytes {
		if bytes[i] != '\r' {
			endIndex++
		} else {
			break
		}
	}
	bulkStringByteLength, err := strconv.Atoi(string(bytes[1:endIndex]))
	if err != nil {
		panic("error reading bulk string length")
	}
	bulkString := string(bytes[endIndex+2 : endIndex+2+bulkStringByteLength])

	return bulkString, endIndex - 1
}

func parseArray(bytes []byte) []string {
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

	results := []string{}
	arrayIndex := endIndex + 2
	for range arrayLength {
		bulkString, elementCount := parseBulkString(bytes[arrayIndex:])
		results = append(results, bulkString)
		arrayIndex += elementCount + 2 + len(bulkString) + 2 + 1
	}

	return results
}
