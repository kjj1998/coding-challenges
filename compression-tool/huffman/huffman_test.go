package huffman

import (
	"testing"
)

func TestHuffmanMethods(t *testing.T) {
	t.Run("Correct huffman tree is built from character frequencies", func(t *testing.T) {
		frequencies := map[rune]int{'C': 32, 'D': 42, 'E': 120, 'K': 7, 'L': 42, 'M': 24, 'U': 37, 'Z': 2}
		gotHuffTree := buildHuffmanTree(frequencies)
		wantHuffTreeRootNodeWeight := 306
		wantHuffTreeRootNodeLeafStatus := false

		if gotHuffTree.Weight() != wantHuffTreeRootNodeWeight {
			t.Errorf("got %d want %d", gotHuffTree.Weight(), wantHuffTreeRootNodeWeight)
		}

		if gotHuffTree.Root().IsLeaf() != wantHuffTreeRootNodeLeafStatus {
			t.Errorf("got %v want %v", gotHuffTree.Root().IsLeaf(), wantHuffTreeRootNodeLeafStatus)
		}
	})

	t.Run("Correct huffman table is built from huffman tree", func(t *testing.T) {
		frequencies := map[rune]int{'C': 32, 'D': 42, 'E': 120, 'K': 7, 'L': 42, 'M': 24, 'U': 37, 'Z': 2}
		huffTree := buildHuffmanTree(frequencies)
		gotHuffTable := buildHuffmanTable(huffTree)

		wantCharacter := 'Z'
		wantCharacterFreq := 2
		var wantCharacterCode uint64 = 0b111100
		wantCharacterBits := 6

		if gotHuffTable[wantCharacter].Freq != wantCharacterFreq {
			t.Errorf("got %d want %d", gotHuffTable[wantCharacter].Freq, wantCharacterFreq)
		}

		if gotHuffTable[wantCharacter].Code != wantCharacterCode {
			t.Errorf("got %b want %b", gotHuffTable[wantCharacter].Code, wantCharacterCode)
		}

		if gotHuffTable[wantCharacter].Bits != wantCharacterBits {
			t.Errorf("got %d want %d", gotHuffTable[wantCharacter].Bits, wantCharacterBits)
		}
	})

	t.Run("Encodes simple text correctly", func(t *testing.T) {
		// Simple example: 'A' -> 0, 'B' ->
		huffTable := map[rune]HuffCode{
			'A': {Code: 0b0, Bits: 1, Freq: 2},
			'B': {Code: 0b1, Bits: 1, Freq: 1},
		}
		text := "AAB"
		// Expected: bits 001 + padding = 00100000 = 0x20

		got := encodeData(text, huffTable)
		want := []byte{0x20}

		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %08b want %08b", got, want)
		}
	})
}
