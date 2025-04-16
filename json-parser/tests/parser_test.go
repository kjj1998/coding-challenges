package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/parser"
)

func TestJsonParser(t *testing.T) {
	t.Run("test step1/valid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step1/valid.json")
		tokens := lexer.Lex(data)

		result := parser.ParseValue(&tokens)
		expected := map[string]interface{}{}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step1/invalid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step1/invalid.json")
		tokens := lexer.Lex(data)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		parser.ParseValue(&tokens)
	})
}
