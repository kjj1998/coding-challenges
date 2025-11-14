package main

import (
	"fmt"

	fileparsing "github.com/kjj1998/coding-challenges/compression-tool/fileparsing"
	huffman "github.com/kjj1998/coding-challenges/compression-tool/huffman"
)

func main() {
	text, err := fileparsing.ParseFile("./files/les-mis.txt")

	if err != nil {
		panic(err)
	}

	// compress and encode file
	fmt.Println("Compressing and encoding file...")
	huffman.CompressFile(text, "./files/encoded/les-mis-encoded")

	// decompress and decode file
	fmt.Println("\nDecompressing and decoding file...")
	huffman.DecompressFile("./files/encoded/les-mis-encoded", "./files/decoded/les-mis-decoded.txt")
}
