package main

import (
	"fmt"
	"os"

	lexer "github.com/kjj1998/coding-challenges/json-parser/lexer"
	parser "github.com/kjj1998/coding-challenges/json-parser/parser"
)

var JSON_LEFTBRACE = "{"
var JSON_LEFTBRACKET = "["
var JSON_RIGHTBRACE = "}"

func main() {
	data, _ := os.ReadFile("./tests/step2/valid.json")

	tokens := lexer.Lex(data)
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
