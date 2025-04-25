package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/models"
)

func Parse(tokens []lexer.Token) (any, error) {
	value, err := ParseValue(&tokens)
	if err != nil {
		return nil, err
	}

	// Enforce JSON root to be either object or array
	switch value.(type) {
	case map[string]any, []any:
		// Good
	default:
		return nil, errors.New("a JSON payload should be an object or array, not a " + getTypeName(value))
	}

	if len(tokens) > 0 {
		return nil, errors.New("unexpected extra tokens after top-level value")
	}

	return value, nil
}

func getTypeName(value any) string {
	switch value.(type) {
	case string:
		return "string"
	case float64, int:
		return "number"
	case bool:
		return "boolean"
	case nil:
		return "null"
	default:
		return "unknown"
	}
}

func ParseValue(tokens *[]lexer.Token) (any, error) {
	if len(*tokens) == 0 {
		panic("unexpected end of input")
	}

	t := (*tokens)[0]
	fmt.Printf("token: %s\n", t.Value)

	switch t.Type {
	case models.LEFT_BRACE:
		obj, err := parseObject(tokens)
		if err != nil {
			return nil, err
		}
		return obj, nil
	case models.LEFT_BRACKET:
		arr, err := parseArray(tokens)
		if err != nil {
			return nil, err
		}
		return arr, nil
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
			num := parseNumber(&t)
			return num, nil
		} else {
			panic("unexpected token: " + t.Value)
		}
	}
}

func parseNumber(token *lexer.Token) any {
	intVal, err := strconv.Atoi(token.Value)
	if err == nil {
		return intVal
	}

	floatVal, err := strconv.ParseFloat(token.Value, 64)
	if err == nil {
		return floatVal
	}

	return nil
}

func parseObject(tokens *[]lexer.Token) (map[string]any, error) {
	obj := make(map[string]any)
	fmt.Printf("token: %s\n", (*tokens)[0].Value)

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

func parseArray(tokens *[]lexer.Token) ([]any, error) {
	slice := []any{}

	consume(tokens, models.LEFT_BRACKET)

	for len(*tokens) > 0 && (*tokens)[0].Type != models.RIGHT_BRACKET {
		value, err := ParseValue(tokens)
		if err != nil {
			return nil, err
		}
		slice = append(slice, value)

		if len(*tokens) > 0 && (*tokens)[0].Type == models.COMMA {
			consume(tokens, models.COMMA)
		} else if len(*tokens) == 0 {
			return nil, errors.New("unclosed array")
		}
	}

	consume(tokens, models.RIGHT_BRACKET)
	return slice, nil
}

func decodeJSONString(input string) (string, error) {
	if len(input) < 2 || input[0] != '"' || input[len(input)-1] != '"' {
		return "", errors.New("invalid JSON string format")
	}
	input = input[1 : len(input)-1] // remove quotes

	var builder strings.Builder
	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch != '\\' {
			builder.WriteByte(ch)
			continue
		}

		// escape sequence starts
		i++
		if i >= len(input) {
			return "", errors.New("incomplete escape sequence")
		}

		switch input[i] {
		case '"':
			builder.WriteByte('"')
		case '\\':
			builder.WriteByte('\\')
		case '/':
			builder.WriteByte('/')
		case 'b':
			builder.WriteByte('\b')
		case 'f':
			builder.WriteByte('\f')
		case 'n':
			builder.WriteByte('\n')
		case 'r':
			builder.WriteByte('\r')
		case 't':
			builder.WriteByte('\t')
		case 'u':
			// Unicode escape: \uXXXX
			if i+4 >= len(input) {
				return "", errors.New("incomplete unicode escape")
			}
			hex := input[i+1 : i+5]
			r, err := strconv.ParseUint(hex, 16, 16)
			if err != nil {
				return "", errors.New("invalid unicode escape: \\u" + hex)
			}
			i += 4

			// Handle UTF-16 surrogate pairs
			r1 := rune(r)
			if utf16.IsSurrogate(r1) && i+6 < len(input) && input[i+1] == '\\' && input[i+2] == 'u' {
				hex2 := input[i+3 : i+7]
				r2, err := strconv.ParseUint(hex2, 16, 16)
				if err == nil {
					i += 6
					combined := utf16.DecodeRune(r1, rune(r2))
					if combined != utf8.RuneError {
						builder.WriteRune(combined)
						continue
					}
				}
			}
			builder.WriteRune(r1)
		default:
			return "", errors.New("unknown escape sequence: \\" + string(input[i]))
		}
	}

	return builder.String(), nil
}

func parseString(tokens *[]lexer.Token) (string, error) {
	t := (*tokens)[0]
	fmt.Printf("string token: %s\n", t.Value)
	consume(tokens, models.STRING)

	decoded, err := decodeJSONString(t.Value)
	if err != nil {
		return "", err
	}
	return decoded, nil
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
