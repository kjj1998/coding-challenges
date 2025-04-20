package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/parser"
)

func TestConclusive(t *testing.T) {
	t.Run("test test_suite/pass1.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/passtest.json")
		tokens, _ := lexer.Lex(data)

		for _, v := range tokens {
			fmt.Println(v)
		}

		result, _ := parser.ParseValue(&tokens)
		expected := []any{
			"JSON Test Pattern pass1",
			map[string]any{
				"object with 1 member": []any{"array with 1 element"},
			},
			map[string]any{},
			[]any{},
			float64(-42),
			true,
			false,
			nil,
			map[string]any{
				"integer": float64(1234567890),
				"real":    -9876.543210,
				"e":       0.123456789e-12,
				"E":       1.234567890e+34,
				"":        23456789012e66,
			},
		}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})
}
