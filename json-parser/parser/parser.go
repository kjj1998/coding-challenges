package parser

import (
	"fmt"
	"strings"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/models"
)

func ParseValue(tokens *[]lexer.Token) any {
	if len(*tokens) == 0 {
		panic("unexpected end of input")
	}

	t := (*tokens)[0]

	switch t.Type {
	case models.LEFT_BRACE:
		return parseObject(tokens)
	default:
		if strings.HasPrefix(t.Value, "\"") {
			return parseString(tokens)
		} else {
			panic("unexpected token: " + t.Value)
		}
	}
}

func parseObject(tokens *[]lexer.Token) map[string]any {
	obj := make(map[string]any)

	consume(tokens, models.LEFT_BRACE)

	for len(*tokens) > 0 && (*tokens)[0].Type != models.RIGHT_BRACE {
		key := parseString(tokens)
		consume(tokens, models.COLON)
		value := ParseValue(tokens)
		obj[key] = value

		if (*tokens)[0].Type == models.COMMA {
			consume(tokens, models.COMMA)
			if (*tokens)[0].Type != models.STRING {
				panic("expected string after comma in object")
			}
		} else {
			break
		}
	}

	consume(tokens, models.RIGHT_BRACE)
	return obj
}

func parseString(tokens *[]lexer.Token) string {
	t := (*tokens)[0]
	consume(tokens, models.STRING)
	fmt.Printf("%s\n", (*tokens)[0].Value)

	return t.Value
}

func consume(tokens *[]lexer.Token, expected models.TokenType) {
	t := (*tokens)[0]
	if len(*tokens) == 0 || t.Type != expected {
		panic("unexpected token: " + t.Value)
	}

	*tokens = (*tokens)[1:]
}
