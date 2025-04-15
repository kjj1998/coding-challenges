package main

import (
	"fmt"
	"os"

	lexer "github.com/kjj1998/coding-challenges/json-parser/lexer"
)

var JSON_LEFTBRACE = "{"
var JSON_LEFTBRACKET = "["

func main() {
	data, _ := os.ReadFile("./tests/step1/valid.json")

	tokens := lexer.Lex(data)

	jsonObject := Parse(&tokens)

	if jsonObject == nil {
		fmt.Println("invalid json")
		os.Exit(1)
	} else {
		fmt.Println("valid json")
		fmt.Println(jsonObject)
		os.Exit(0)
	}
}

func Parse(tokens *[]string) map[string]any {
	obj := make(map[string]any)

	if len(*tokens) == 0 {
		return nil
	}

	t := (*tokens)[0]

	if t == JSON_LEFTBRACE {
		*tokens = (*tokens)[1:]
	}

	for len(*tokens) > 0 && (*tokens)[0] != "}" {

	}

	if t == "}" {
		*tokens = (*tokens)[1:]
	}
	return obj
}

// func parseObject(tokens *[]string) map[string]interface{} {
// 	jsonObject := make(map[string]interface{})
// 	t := (*tokens)[0]

// 	if t == "}" {
// 		*tokens = (*tokens)[1:]
// 		return jsonObject
// 	}

// 	for {
// 		jsonKey := (*tokens)[0]
// 		*tokens = (*tokens)[1:]

// 		*tokens = (*tokens)[1:]
// 		jsonValue := Parse(tokens)

// 		jsonObject[jsonKey] = jsonValue

// 		t = (*tokens)[0]
// 		if t == "}" {
// 			return jsonObject
// 		}

// 		*tokens = (*tokens)[1:]
// 	}
// }
