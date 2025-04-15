package lexer

var JSON_SYNTAX = map[rune]struct{}{
	'{': {},
	'}': {},
	':': {},
}

var JSON_WHITESPACE = map[rune]struct{}{
	' ':  {},
	'\n': {},
}

func Lex(data []byte) []string {
	var tokens []string

	for len(data) != 0 {
		json_string := lex_string(&data)

		if json_string != "" {
			tokens = append(tokens, json_string)
			continue
		}

		if _, ok := JSON_SYNTAX[rune(data[0])]; ok {
			tokens = append(tokens, string(data[0]))
			data = data[1:]
		} else if _, ok := JSON_WHITESPACE[rune(data[0])]; ok {
			data = data[1:]
		}
	}

	return tokens
}

func lex_string(data *[]byte) string {
	var runes []rune
	slice := *data

	if slice[0] == '"' {
		slice = slice[1:]
	} else {
		return ""
	}

	for _, c := range slice {

		if c == '"' {
			slice = slice[len(runes)+1:]
			*data = slice
			return string(runes)
		} else {
			runes = append(runes, rune(c))
		}
	}

	*data = slice
	return string(runes)
}
