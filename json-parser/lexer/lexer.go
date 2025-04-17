package lexer

import (
	"errors"
	"strings"
	"unicode"

	"github.com/kjj1998/coding-challenges/json-parser/models"
)

type Token = models.Token

var JSON_SYNTAX = map[rune]struct{}{
	'{': {},
	'}': {},
	':': {},
}

var JSON_WHITESPACE = map[rune]struct{}{
	' ':  {},
	'\n': {},
}

func Lex(data []byte) ([]Token, error) {
	tokens := []Token{}
	i := 0

	for i < len(data) {
		c := data[i]

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
			str, consumed := lexString(data[i:])
			tokens = append(tokens, Token{Type: models.STRING, Value: str})
			i += consumed
		default:
			if isDigit(c) || c == '-' {
				num, consumed := lexNumber(data)
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
			} else {
				return tokens, errors.New("unexpected character: " + string(c))
			}
		}
	}

	return tokens, nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func lexNumber(data []byte) (string, int) {
	i := 0
	for i < len(data) && (unicode.IsDigit(rune(data[i])) || data[i] == '.' || data[i] == '-' || data[i] == 'e' || data[i] == 'E' || data[i] == '+') {
		i++
	}

	return string(data[:i]), i
}

func lexString(data []byte) (string, int) {
	if len(data) == 0 || data[0] != '"' {
		panic("string must start with '\"'")
	}

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
		panic("unterminated string")
	}

	raw := string(data[:i+1])
	return raw, i + 1
}
