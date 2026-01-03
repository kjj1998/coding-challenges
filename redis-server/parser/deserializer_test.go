package parser

import (
	"bufio"
	"bytes"
	"errors"
	"testing"
)

func TestDeserialize(t *testing.T) {
	/* Simple String */
	// deserialize simple string
	t.Run("Deserialize +OK\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("+OK\r\n")

		got, _ := Deserialize(serializedCommand.Bytes())
		want := []string{"OK"}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize simple string
	t.Run("Deserialize +OK\r\n", func(t *testing.T) {
		data := []byte("+OK\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		got, _ := ReadSimpleString(bufio.NewReader(reader))
		want := []byte("OK")

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize simple string
	t.Run("Deserialize +hello world\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("+hello world\r\n")

		got, _ := Deserialize(serializedCommand.Bytes())
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

		got, _ := Deserialize(serializedCommand.Bytes())
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

		_, gotErr := Deserialize(serializedCommand.Bytes())
		wantErr := "Error message must begin with ERR: unknown command"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// deserialize integer
	t.Run("Deserialize :1000\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString(":1000\r\n")

		got, _ := Deserialize(serializedCommand.Bytes())
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

		_, gotErr := Deserialize(serializedCommand.Bytes())
		wantErr := "non integer input: abcde"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// deserialize bulk string
	t.Run("Deserialize $5\r\nhello\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$5\r\nhello\r\n")

		got, _ := Deserialize(serializedCommand.Bytes())
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

		_, gotErr := Deserialize(serializedCommand.Bytes())
		wantErr := "non integer value given for bulk string length: abcde"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// bulk string length input greater than length of bulk string
	t.Run("Deserialize $10\r\nhello\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$10\r\nhello\r\n")

		_, gotErr := Deserialize(serializedCommand.Bytes())
		wantErr := "incorrect bulk string length input: 10"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	// bulk string length input lesser than length of bulk string
	t.Run("Deserialize $3\r\nhelloworld\r\n", func(t *testing.T) {
		serializedCommand := bytes.NewBufferString("$3\r\nhelloworld\r\n")

		_, gotErr := Deserialize(serializedCommand.Bytes())
		wantErr := "incorrect bulk string length input: 3"

		if gotErr.Error() != wantErr {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr)
		}
	})

	/* ARRAYS */
	// deserialize array with two bulk strings
	t.Run("Deserialize *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", func(t *testing.T) {
		data := []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		got, _ := DeserializeArray(bufio.NewReader(reader))
		want := [][]byte{[]byte("hello"), []byte("world")}

		for i := range got {
			if string(got[i].([]byte)) != string(want[i]) {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize array with mixed types for elements
	t.Run("Deserialize *3\r\n$5\r\nhello\r\n:555\r\n$5\r\nworld\r\n", func(t *testing.T) {
		data := []byte("*3\r\n$5\r\nhello\r\n:555\r\n$5\r\nworld\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		got, _ := DeserializeArray(bufio.NewReader(reader))
		want := [][]byte{[]byte("hello"), []byte("555"), []byte("world")}

		for i := range got {
			if string(got[i].([]byte)) != string(want[i]) {
				t.Errorf("got %q want %q", got[i], want[i])
			}
		}
	})

	// deserialize array with no elements
	t.Run("Deserialize *0\r\n", func(t *testing.T) {
		data := []byte("*0\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		got, _ := DeserializeArray(bufio.NewReader(reader))
		want := [][]byte{}

		if len(got) != len(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// deserialize array with nested arrays within
	t.Run("Deserialize nested array *3\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n:555\r\n$5\r\nworld\r\n", func(t *testing.T) {
		data := []byte("*3\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n:555\r\n$5\r\nworld\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		got, _ := DeserializeArray(bufio.NewReader(reader))
		want := []any{[][]byte{[]byte("hello"), []byte("world")}, []byte("555"), []byte("world")}

		for i := range got {
			if i == 0 {
				gotArr := got[i].([]any)
				wantArr := want[i].([][]byte)
				for j := range gotArr {
					gotString := string(gotArr[j].([]byte))
					wantString := string(wantArr[j])
					if gotString != wantString {
						t.Errorf("got %q want %q", gotString, wantString)
					}
				}
			} else {
				gotString := string(got[i].([]byte))
				wantString := string(want[i].([]byte))
				if gotString != wantString {
					t.Errorf("got %q want %q", gotString, wantString)
				}
			}
		}
	})

	// deserialize array where the length of the array given is lesser than the actual number of elements
	t.Run("Deserialize *1\r\n$3\r\nget\r\n$3\r\nkey\r\n", func(t *testing.T) {
		data := []byte("*1\r\n$3\r\nget\r\n$3\r\nkey\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		_, gotErr := DeserializeArray(bufio.NewReader(reader))
		wantErr := errors.New("incorrect number of elements given: 1")

		if gotErr.Error() != wantErr.Error() {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr.Error())
		}
	})

	// deserialize array where the length of the array given is more than the actual number of elements
	t.Run("Deserialize *4\r\n$3\r\nget\r\n$3\r\nkey\r\n", func(t *testing.T) {
		data := []byte("*4\r\n$3\r\nget\r\n$3\r\nkey\r\n")
		reader := bytes.NewReader(data)
		reader.ReadByte()

		_, gotErr := DeserializeArray(bufio.NewReader(reader))
		wantErr := errors.New("incorrect number of elements given: 4")

		if gotErr.Error() != wantErr.Error() {
			t.Errorf("got %q want %q", gotErr.Error(), wantErr.Error())
		}
	})
}
