package utils

import (
	"bufio"
	"log"
	"os"
)

func Read(filePath *string) *[]string {
	var scanner *bufio.Scanner

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

	strs := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		strs = append(strs, line)
	}

	return &strs
}
