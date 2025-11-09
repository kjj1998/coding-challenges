package main

import (
	"fmt"

	fileparsing "github.com/kjj1998/coding-challenges/compression-tool/fileparsing"
	huffman "github.com/kjj1998/coding-challenges/compression-tool/huffman"
)

func main() {
	fileparsing.ParseFile("./files/test.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// for key, val := range data {
	// 	fmt.Printf("Key: %c, Value: %d\n", key, val)
	// }

	// Some items and their priorities.
	items := map[rune]int{
		'C': 32, 'D': 42, 'E': 120, 'K': 7, 'L': 42, 'M': 24, 'U': 37, 'Z': 2,
	}

	huffTree := huffman.BuildHuffmanTree(items)
	huffTable := huffman.BuildHuffmanTable(huffTree)

	for k, v := range huffTable {
		fmt.Printf("key = %c, freq = %d, code = %b, bits = %d\n", k, v.Freq, v.Code, v.Bits)
	}
}
