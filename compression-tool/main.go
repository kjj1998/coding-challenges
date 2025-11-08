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

	tree := huffman.BuildHuffmanTree(items)

	tree.PreOrderTraversal(func(node huffman.HuffBaseNode, code byte) {
		if node.IsLeaf() {
			leaf := node.(*huffman.HuffLeafNode)
			fmt.Printf("Leaf element: %c (weight: %d, code: %b)\n", leaf.Element(), leaf.Weight(), code)
		} else {
			fmt.Printf("Internal (weight: %d)\n", node.Weight())
		}
	}, 0)
}
