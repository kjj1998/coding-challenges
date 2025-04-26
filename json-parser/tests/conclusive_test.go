package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/kjj1998/coding-challenges/json-parser/lexer"
	"github.com/kjj1998/coding-challenges/json-parser/parser"
)

func TestConclusive(t *testing.T) {
	t.Run("test test_suite/pass1.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/pass1.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := []any{
			"JSON Test Pattern pass1",
			map[string]any{
				"object with 1 member": []any{"array with 1 element"},
			},
			map[string]any{},
			[]any{},
			-42,
			true,
			false,
			nil,
			map[string]any{
				"integer":       1234567890,
				"real":          -9876.543210,
				"e":             0.123456789e-12,
				"E":             1.234567890e+34,
				"":              23456789012e66,
				"zero":          0,
				"one":           1,
				"space":         " ",
				"quote":         "\"",
				"backslash":     "\\",
				"controls":      "\b\f\n\r\t",
				"slash":         "/ & /",
				"alpha":         "abcdefghijklmnopqrstuvwyz",
				"ALPHA":         "ABCDEFGHIJKLMNOPQRSTUVWYZ",
				"digit":         "0123456789",
				"0123456789":    "digit",
				"special":       "`1~!@#$%^&*()_+-={':[,]}|;.</>?",
				"hex":           "\u0123\u4567\u89AB\uCDEF\uabcd\uef4A",
				"true":          true,
				"false":         false,
				"null":          nil,
				"array":         []any{},
				"object":        map[string]any{},
				"address":       "50 St. James Street",
				"url":           "http://www.JSON.org/",
				"comment":       "// /* <!-- --",
				"# -- --> */":   " ",
				" s p a c e d ": []any{1, 2, 3, 4, 5, 6, 7},
				"compact":       []any{1, 2, 3, 4, 5, 6, 7},
				"jsontext":      "{\"object with 1 member\":[\"array with 1 element\"]}",
				"quotes":        "&#34; \u0022 %22 0x22 034 &#x22;",
				"/\\\"\uCAFE\uBABE\uAB98\uFCDE\ubcda\uef4A\b\f\n\r\t`1~!@#$%^&*()_+-=[]{}|;:',./<>?": "A key can be any string",
			},
			0.5, 98.6, 99.44,
			1066,
			1e1,
			0.1e1,
			1e-1,
			1e00, 2e+00, 2e-00,
			"rosebud",
		}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test test_suite/pass2.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/pass2.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := []any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{[]any{"Not too deep"}}}}}}}}}}}}}}}}}}}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test test_suite/pass3.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/pass3.json")
		tokens, _ := lexer.Lex(data)

		result, _ := parser.ParseValue(&tokens)
		expected := map[string]any{
			"JSON Test Pattern pass3": map[string]any{
				"The outermost value": "must be an object or array.",
				"In this test":        "It is an object.",
			},
		}

		if result == nil {
			t.Error("nil not expected")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v\n", expected, result)
		}
	})

	t.Run("test test_suite/fail1.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail1.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "a JSON payload should be an object or array, not a string"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail2.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail2.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unclosed array"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail3.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail3.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "invalid JSON string format"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail4.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail4.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "extra comma in array"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail5.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail5.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected comma in array"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail6.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail6.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected comma in array"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail7.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail7.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected extra tokens after top-level value"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail8.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail8.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected extra tokens after top-level value"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail9.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail9.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "expected string after comma in object"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail10.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail10.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected extra tokens after top-level value"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail11.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail11.json")
		_, err := lexer.Lex(data)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected character: +"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail12.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail12.json")
		_, err := lexer.Lex(data)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unterminated string"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail13.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail13.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected token: 013"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail14.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail14.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unexpected end of input"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail15.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail15.json")
		tokens, _ := lexer.Lex(data)

		_, err := parser.Parse(tokens)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		expected := "unknown escape sequence: \\x"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail16.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail16.json")
		_, err := lexer.Lex(data)

		expected := "unexpected character: \\"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("test test_suite/fail17.json", func(t *testing.T) {
		data, _ := os.ReadFile("./test_suite/fail17.json")
		tokens, _ := lexer.Lex(data)
		_, err := parser.Parse(tokens)

		expected := "unknown escape sequence: \\0"
		if err.Error() != expected {
			t.Errorf("expected error %q, got %q", expected, err.Error())
		}
	})
}
