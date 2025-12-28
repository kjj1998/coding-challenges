package parser

import (
	"bytes"
	"testing"
)

func TestDeserialize(t *testing.T) {
	t.Run("Deserialize +OK\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("+OK\r\n")

		got, _ := Deserialize(*serializedCommand)
		want := []string{"OK"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	t.Run("Deserialize +hello world\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("+hello world\r\n")

		got, _ := Deserialize(*serializedCommand)
		want := []string{"hello world"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	t.Run("Deserialize -ERR unknown command\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("-ERR unknown command\r\n")

		got, _ := Deserialize(*serializedCommand)
		want := []string{"ERR unknown command"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	t.Run("Deserialize :1000\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString(":1000\r\n")

		got, _ := Deserialize(*serializedCommand)
		want := []string{"1000"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	t.Run("Deserialize $5\r\nhello\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$5\r\nhello\r\n")

		got, _ := Deserialize(*serializedCommand)
		want := []string{"hello"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	t.Run("Deserialize *2\r\n$5\r\nhello\r\n$3\r\nbar\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("*2\r\n$5\r\nhello\r\n$3\r\nbar\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"hello", "bar"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})
}
