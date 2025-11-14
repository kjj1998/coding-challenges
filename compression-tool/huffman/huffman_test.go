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
}
