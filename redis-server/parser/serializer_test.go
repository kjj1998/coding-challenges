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
}
