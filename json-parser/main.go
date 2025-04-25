package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/parser"
)

func main() {
	data, _ := os.ReadFile("./tests/step4/valid.json")

	tokens, err := lexer.Lex(data)
	fmt.Println(err)
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
