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

	t.Run("parse file and get the contents of the text", func(t *testing.T) {
		got, _ := ParseFile("../files/test.txt")
		want := "\ufeffThe Project Gutenberg eBook of Les Misérables, by Victor Hugo\r\n\r\nThis eBook is for the use of anyone anywhere in the United States and\r\nmost other parts of the world at no cost and with almost no restrictions\r\nwhatsoever. You may copy it, give it away or re-use it under the terms\r\nof the Project Gutenberg License included with this eBook or online at\r\nwww.gutenberg.org. If you are not located in the United States, you\r\nwill have to check the laws of the country where you are located before\r\nusing this eBook. Hello World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
