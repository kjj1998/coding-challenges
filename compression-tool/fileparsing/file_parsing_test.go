package fileparsing

import (
	"fmt"
	"testing"
)

func TestParsefile(t *testing.T) {
	t.Run("file does not exist", func(t *testing.T) {
		_, got := ParseFile("../files/test123.txt")
		want := fmt.Errorf("../files/test123.txt: does not exist")

		if got.Error() != want.Error() {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("no permission to open file", func(t *testing.T) {
		_, got := ParseFile("../files/permission-denied.txt")
		want := fmt.Errorf("../files/permission-denied.txt: permission denied")

		if got.Error() != want.Error() {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("parse file and get frequency of the X character", func(t *testing.T) {
		got, _ := ParseFile("../files/test.txt")
		want := 333

		if got['X'] != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("parse file and get frequency of the t character", func(t *testing.T) {
		got, _ := ParseFile("../files/test.txt")
		want := 223000

		if got['t'] != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
