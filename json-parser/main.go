package main

import (
	"fmt"
	"os"

	lexer "github.com/kjj1998/coding-challenges/json-parser/lexer"
)

func main() {
	data, _ := os.ReadFile("./tests/step2/valid.json")

	tokens := lexer.Lex(data)

	for _, token := range tokens {
		fmt.Println(token)
	}
}
