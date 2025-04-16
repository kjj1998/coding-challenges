package models

type TokenType int

const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	COLON
	COMMA
	STRING
	NUMBER
	TRUE
	FALSE
	NULL
)

type Token struct {
	Type  TokenType
	Value string
}
