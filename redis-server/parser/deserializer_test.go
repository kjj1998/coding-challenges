package parser

import (
	"bytes"
	"testing"
)

func TestDeserialize(t *testing.T) {
	// deserialize simple string
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

	// deserialize simple string
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

	// deserialize error message
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

	// error message not starting with ERR
	t.Run("Deserialize -unknown command\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("-unknown command\r\n")

		_, gotErr := Deserialize(*serializedCommand)
		wantErr := "Error message must begin with ERR: unknown command"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// deserialize integer
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

	// non integer supplied
	t.Run("Deserialize :abcde\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString(":abcde\r\n")

		_, gotErr := Deserialize(*serializedCommand)
		wantErr := "non integer input: abcde"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// deserialize bulk string
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

	// non integer bulk string length input
	t.Run("Deserialize $abcde\r\nhello\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$abcde\r\nhello\r\n")

		_, gotErr := Deserialize(*serializedCommand)
		wantErr := "non integer value given for bulk string length: abcde"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// bulk string length input greater than length of bulk string
	t.Run("Deserialize $10\r\nhello\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$10\r\nhello\r\n")

		_, gotErr := Deserialize(*serializedCommand)
		wantErr := "incorrect bulk string length input: 10"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// bulk string length input lesser than length of bulk string
	t.Run("Deserialize $3\r\nhelloworld\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$3\r\nhelloworld\r\n")

		_, gotErr := Deserialize(*serializedCommand)
		wantErr := "incorrect bulk string length input: 3"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// deserialize array with two bulk strings
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

	// deserialize array with one bulk string
	t.Run("Deserialize *1\r\n$4\r\nping\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("*1\r\n$4\r\nping\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"ping"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize array with two bulk strings
	t.Run("Deserialize *2\r\n$4\r\necho\r\n$11\r\nhello world\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"echo", "hello world"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize array with two bulk strings
	t.Run("Deserialize *2\r\n$3\r\nget\r\n$3\r\nkey\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("*2\r\n$3\r\nget\r\n$3\r\nkey\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"get", "key"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize bulk string with no string
	t.Run("Deserialize $0\r\n\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$0\r\n\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{""}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize array with no elements
	t.Run("Deserialize $0\r\n\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$0\r\n\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{""}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserizlie null value for bulk string
	t.Run("Deserialize $-1\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$-1\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"null"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize null value for array
	t.Run("Deserialize *-1\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("*-1\r\n")
		got, _ := Deserialize(*serializedCommand)
		want := []string{"null"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})
}
