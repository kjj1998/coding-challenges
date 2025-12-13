package algo

import (
	"slices"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Run("Sort an array of strings", func(t *testing.T) {
		testStr := []string{"banana", "app", "apple", "bat"}

		got := HeapSort(testStr)
		want := []string{"app", "apple", "banana", "bat"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Sort an array of strings with duplicates", func(t *testing.T) {
		testStr := []string{"bat", "banana", "app", "apple", "bat", "app"}

		got := HeapSort(testStr)
		want := []string{"app", "app", "apple", "banana", "bat", "bat"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
