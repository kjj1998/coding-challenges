package main

import (
	"fmt"
	"os"

	lexer "github.com/kjj1998/coding-challenges/json-parser/lexer"
	parser "github.com/kjj1998/coding-challenges/json-parser/parser"
)

func main() {
	data, _ := os.ReadFile("./tests/step2/valid.json")

	tokens := lexer.Lex(data)

	for _, v := range tokens {
		fmt.Println(v)
	}

	jsonObject := parser.ParseValue(&tokens)

	if jsonObject == nil {
		fmt.Println("invalid json")
		os.Exit(1)
	} else {
		fmt.Println("valid json")
		fmt.Println(jsonObject)
		os.Exit(0)
	}
}
