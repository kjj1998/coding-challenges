package cut

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Cut(fieldPositions []int, delimiter *string, filePath *string) [][]string {
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

	for _, position := range fieldPositions {
		if position < 1 || position > len(data) {
			log.Fatalf("Field position value out of bounds")
		}
	}

	var cutValues [][]string

	for _, position := range fieldPositions {
		cutValues = append(cutValues, data[position-1]) // position - 1 because user supplied field positions start as 1 instead of 0
	}

	return cutValues
}

func checkFieldPositions(fieldPositions []string) []int {
	intPositions := []int{}

	for _, position := range fieldPositions {
		intPosition, err := strconv.Atoi(position)
		if err != nil {
			log.Fatal("Only integer values can be entered for field positions\n")
		}
		intPositions = append(intPositions, intPosition)
	}

	return intPositions
}

func ParseFieldPos(fieldPos *string) []int {
	if len(*fieldPos) == 1 {
		return checkFieldPositions([]string{*fieldPos})
	}

	if strings.ContainsRune(*fieldPos, ',') {
		return checkFieldPositions(strings.Split(*fieldPos, ","))
	}

	if strings.ContainsRune(*fieldPos, ' ') {
		return checkFieldPositions(strings.Split(*fieldPos, " "))
	}

	log.Fatalf("Invalid field position(s) entered: %s\n", *fieldPos)
	return nil
}
