package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/kjj1998/coding-challenges/json-parser/models"
)

type Token = models.Token

func Lex(data []byte) ([]Token, error) {
	tokens := []Token{}
	i := 0

	fmt.Printf("length: %d\n", len(data))
	for i < len(data) {
		c := data[i]
		fmt.Printf("character: %c\n", c)

		switch c {
		case ' ', '\n', '\r', '\t': // check for whitespace
			i++
		case '{': // check for left brace
			tokens = append(tokens, Token{Type: models.LEFT_BRACE, Value: "{"})
			i++
		case '}': // check for right brace
			tokens = append(tokens, Token{Type: models.RIGHT_BRACE, Value: "}"})
			i++
		case ':': // check for colon
			tokens = append(tokens, Token{Type: models.COLON, Value: ":"})
			i++
		case ',': // check for comma
			tokens = append(tokens, Token{Type: models.COMMA, Value: ","})
			i++
		case '[': // check for left bracket
			tokens = append(tokens, Token{Type: models.LEFT_BRACKET, Value: "["})
			i++
		case ']': // check for right bracket
			tokens = append(tokens, Token{Type: models.RIGHT_BRACKET, Value: "]"})
			i++
		case '"': // check for string
			str, consumed, err := lexString(data[i:])
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{Type: models.STRING, Value: str})
			i += consumed
		default:
			if isDigit(c) || c == '-' {
				num, consumed := lexNumber(data, i)
				tokens = append(tokens, Token{Type: models.NUMBER, Value: num})
				i += consumed
			} else if strings.HasPrefix(string(data[i:]), "true") {
				tokens = append(tokens, Token{Type: models.TRUE, Value: "true"})
				i += 4
			} else if strings.HasPrefix(string(data[i:]), "false") {
				tokens = append(tokens, Token{Type: models.FALSE, Value: "false"})
				i += 5
			} else if strings.HasPrefix(string(data[i:]), "null") {
				tokens = append(tokens, Token{Type: models.NULL, Value: "null"})
				i += 4
			} else if isIdentifierStart(c) {
				str, consumed, err := lexString(data[i:])
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, Token{Type: models.STRING, Value: str})
				i += consumed
			} else {
				fmt.Println("unexpected character: " + string(c))
				return tokens, errors.New("unexpected character: " + string(c))
			}
		}
	}

	return tokens, nil
}

func isIdentifierStart(c byte) bool {
	return unicode.IsLetter(rune(c)) || c == '_'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func lexNumber(data []byte, i int) (string, int) {
	start, end := i, i
	for end < len(data) && (unicode.IsDigit(rune(data[end])) || data[end] == '.' || data[end] == '-' || data[end] == 'e' || data[end] == 'E' || data[end] == '+') {
		end++
	}

	return string(data[start:end]), end - start
}

func lexString(data []byte) (string, int, error) {
	// if len(data) == 0 || data[0] != '"' {
	// 	panic("string must start with '\"'")
	// }

	i := 1 // skip opening quote

	for i < len(data) {
		if data[i] == '"' {
			// Count backslashes before the quote
			backslashCount := 0
			j := i - 1
			for j >= 0 && data[j] == '\\' {
				backslashCount++
				j--
			}

			// if backslash count is even, quote is not escaped
			if backslashCount%2 == 0 {
				break
			}
		}

		i++
	}

	if i >= len(data) {
		return "", 0, errors.New("unterminated string")
	}

	raw := string(data[:i+1])
	return raw, i + 1, nil
}
