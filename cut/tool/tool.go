package tool

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Cut(fieldPos *string, delimiter *string, filePath *string) []string {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitValues := strings.Split(line, *delimiter)

		for i, val := range splitValues {
			if len(data) < i+1 {
				data = append(data, []string{val})
			} else {
				data[i] = append(data[i], val)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fieldPosInteger, err := strconv.Atoi(*fieldPos)
	if err != nil {
		log.Fatalf("Error converting %s to integer: %v", *fieldPos, err)
	}

	if fieldPosInteger < 1 || fieldPosInteger > len(data) {
		log.Fatalf("Field position value out of bounds")
	}

	return data[fieldPosInteger-1]
}
