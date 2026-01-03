package parser

import "testing"

func TestSerialize(t *testing.T) {
	// Serialize OK as simple string
	t.Run("Serialize OK to +OK\r\n", func(t *testing.T) {
		got := string(SerializeSimpleString("OK"))
		want := "+OK\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// Serialize PING OK as simple string
	t.Run("Serialize PING OK", func(t *testing.T) {
		got := string(SerializeSimpleString("PING OK"))
		want := "+PING OK\r\n"

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

	// Serialize ERR
	t.Run("Serialize Generic Error to -ERR Generic Error\r\n", func(t *testing.T) {
		got := string(SerializeError("Generic Error"))
		want := "-ERR Generic Error\r\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
