package main

import (
	fileparsing "github.com/kjj1998/coding-challenges/compression-tool/fileparsing"
	huffman "github.com/kjj1998/coding-challenges/compression-tool/huffman"
)

func main() {
	text, err := fileparsing.ParseFile("./files/test.txt")

	if err != nil {
		panic(err)
	}

	// compress file
	huffman.CompressFile(text, "./files/encoded/test")

	// decompress file
	huffman.DecompressFile("./files/encoded/test", "./files/decoded/decoded-test.txt")
}
