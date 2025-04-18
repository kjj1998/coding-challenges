package parser

import (
	"errors"
	"strconv"
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
		} else if strings.HasPrefix(t.Value, "true") {
			consume(tokens, models.TRUE)
			return true, nil
		} else if strings.HasPrefix(t.Value, "false") {
			consume(tokens, models.FALSE)
			return false, nil
		} else if strings.HasPrefix(t.Value, "null") {
			consume(tokens, models.NULL)
			return nil, nil
		} else if isNumber(t.Value) {
			consume(tokens, models.NUMBER)
			return strconv.ParseFloat(t.Value, 64)
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

func isNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return true
	}

	return false
}
