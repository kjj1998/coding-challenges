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
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step1/invalid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step1/invalid.json")
		tokens, _ := lexer.Lex(data)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		parser.ParseValue(&tokens)
	})

	t.Run("test step2/invalid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step2/invalid.json")
		tokens, _ := lexer.Lex(data)
		_, err := parser.ParseValue(&tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "expected string after comma in object"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test step2/invalid2.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step2/invalid2.json")
		_, err := lexer.Lex(data)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected character: k"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test step2/valid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step2/valid.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{"key": "value"}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step2/valid2.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step2/valid2.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{"key": "value", "key2": "value"}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step3/valid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step3/valid.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{"key1": true, "key2": false, "key3": nil, "key4": "value", "key5": 101}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step3/invalid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step3/invalid.json")
		_, err := lexer.Lex(data)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected character: F"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test step4/valid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step4/valid.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{"key": "value", "key-l": []any{}, "key-n": 101, "key-o": map[string]any{}}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step4/valid2.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step4/valid2.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{"key": "value", "key-l": []any{"list value"}, "key-n": 101, "key-o": map[string]any{"inner key": "inner value"}}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test step4/invalid.json", func(t *testing.T) {
		data, _ := os.ReadFile("./step4/invalid.json")
		_, err := lexer.Lex(data)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected character: '"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})
}
