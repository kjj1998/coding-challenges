package parser

import (
	"errors"
	"strings"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/models"
)

func ParseValue(tokens *[]lexer.Token) (any, error) {
	if len(*tokens) == 0 {
		panic("unexpected end of input")
	}

	t := (*tokens)[0]

	switch t.Type {
	case models.LEFT_BRACE:
		obj, err := parseObject(tokens)
		if err != nil {
			return nil, err
		}
		return obj, nil
	default:
		if strings.HasPrefix(t.Value, "\"") {
			str, err := parseString(tokens)
			if err != nil {
				return nil, err
			}
			return str, err
		} else {
			panic("unexpected token: " + t.Value)
		}
	}
}

func parseObject(tokens *[]lexer.Token) (map[string]any, error) {
	obj := make(map[string]any)

	consume(tokens, models.LEFT_BRACE)

	for len(*tokens) > 0 && (*tokens)[0].Type != models.RIGHT_BRACE {
		key, err := parseString(tokens)
		if err != nil {
			return nil, err
		}

		consume(tokens, models.COLON)

		value, err := ParseValue(tokens)
		if err != nil {
			return nil, err
		}
		obj[key] = value

		if (*tokens)[0].Type == models.COMMA {
			consume(tokens, models.COMMA)
			if (*tokens)[0].Type != models.STRING {
				return nil, errors.New("expected string after comma in object")
			}
		} else {
			break
		}
	}

	consume(tokens, models.RIGHT_BRACE)
	return obj, nil
}

func parseString(tokens *[]lexer.Token) (string, error) {
	t := (*tokens)[0]
	consume(tokens, models.STRING)

	return t.Value, nil
}

func consume(tokens *[]lexer.Token, expected models.TokenType) {
	t := (*tokens)[0]
	if len(*tokens) == 0 || t.Type != expected {
		panic("unexpected token: " + t.Value)
	}

	*tokens = (*tokens)[1:]
}
