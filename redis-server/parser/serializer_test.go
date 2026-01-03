package parser

import "testing"

func TestSerialize(t *testing.T) {
	t.Run("Serialize PING", func(t *testing.T) {
		testCommand := []string{"PING"}

		got, _ := Serialize(testCommand)
		want := "*1\r\n$4\r\nPING\r\n"

		if got.String() != want {
			t.Errorf("got %q want %q", got.String(), want)
		}
	})

	t.Run("Serialize integer 1000", func(t *testing.T) {
		testCommand := []string{"1000"}

		got, _ := Serialize(testCommand)
		want := "*1\r\n$4\r\n1000\r\n"

		if got.String() != want {
			t.Errorf("got %q want %q", got.String(), want)
		}
	})

	t.Run("Serialize PING OK", func(t *testing.T) {
		testCommand := []string{"PING", "OK"}

		got, _ := Serialize(testCommand)
		want := "*2\r\n$4\r\nPING\r\n$2\r\nOK\r\n"

		if got.String() != want {
			t.Errorf("got %q want %q", got.String(), want)
		}
	})

	// Serialize OK as simple string
	t.Run("Serialize OK to +OK\r\n", func(t *testing.T) {
		got := string(SerializeSimpleString("OK"))
		want := "+OK\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// Serialize hello world as bulk string
	t.Run("Serialize hello world to $11\r\nhello world\r\n", func(t *testing.T) {
		bulkString := []byte("hello world")

		got := string(SerializeBulkString(bulkString))
		want := "$11\r\nhello world\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// Serialize 1000 as integer
	t.Run("Serialize 1000 to :1000\r\n", func(t *testing.T) {
		got := string(SerializeInteger(1000))
		want := ":1000\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// Serialize null
	t.Run("Serialize null to _\r\n", func(t *testing.T) {
		got := string(SerializeNull())
		want := "_\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
