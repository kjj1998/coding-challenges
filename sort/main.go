package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	flag.Parse()

	var filePath string
	if len(flag.Args()) > 0 {
		filePath = flag.Args()[0]
	}

	var scanner *bufio.Scanner

	file, err := os.Open(filePath)
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

	slices.Sort(strs)

	for _, str := range strs {
		fmt.Println(str)
	}
}
