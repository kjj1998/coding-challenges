package main

import (
	"fmt"
	"log"
	"os"

	lexer "github.com/kjj1998/coding-challenges/json-parser/lexer"
	parser "github.com/kjj1998/coding-challenges/json-parser/parser"
)

func main() {
	data, _ := os.ReadFile("./tests/step2/invalid2.json")

	tokens, err := lexer.Lex(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range tokens {
		fmt.Println(v)
	}

	jsonObject, err := parser.ParseValue(&tokens)

	if err != nil {
		log.Fatalf("invalid json: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("valid json: %v\n", jsonObject)
	os.Exit(0)
}
