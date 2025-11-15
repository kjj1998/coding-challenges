package huffman

import (
	"container/heap"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

/* PriorityQueue implementation */

type PriorityQueue []HuffBaseNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight() < pq[j].Weight()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(HuffBaseNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

/* Compression and encoding functions */

func buildHuffmanTree(frequencies map[rune]int) *HuffTree {
	pq := make(PriorityQueue, 0, len(frequencies))
	for char, freq := range frequencies {
		leaf := &HuffLeafNode{
			element: char,
			weight:  freq,
		}

		pq = append(pq, leaf)
	}

	heap.Init(&pq)
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(HuffBaseNode)
		right := heap.Pop(&pq).(HuffBaseNode)

		internal := &HuffIntermediateNode{
			left:   left,
			right:  right,
			weight: left.Weight() + right.Weight(),
		}

		heap.Push(&pq, internal)
	}
	root := heap.Pop(&pq).(HuffBaseNode)

	return &HuffTree{root: root}
}

func buildHuffmanTable(huffTree *HuffTree) map[rune]HuffCode {
	huffTable := huffTree.PreOrderTraversal(func(node HuffBaseNode, code uint64, bits int, huffTable map[rune]HuffCode) {
		if node.IsLeaf() {
			leaf := node.(*HuffLeafNode)
			huffCode := HuffCode{Freq: leaf.Weight(), Code: code, Bits: bits}
			huffTable[leaf.Element()] = huffCode
		}
	})

	return huffTable
}

func encodeData(text string, huffTable map[rune]HuffCode) []byte {
	var result []byte
	var currentByte byte = 0
	var bitPosition int = 0

	for _, char := range text {
		code := huffTable[char]

		// Write each bit of the code
		for i := code.Bits - 1; i >= 0; i-- {
			bit := byte((code.Code >> i) & 1)
			currentByte |= (bit << (7 - bitPosition))
			bitPosition++

			if bitPosition == 8 {
				result = append(result, currentByte)
				currentByte = 0
				bitPosition = 0
			}
		}
	}

	if bitPosition > 0 {
		result = append(result, currentByte)
	}

	return result
}

func writeEncodedDataToFile(filename string, text string, huffTable map[rune]HuffCode, compressedData []byte) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the Huffman table header
	numChars := uint32(len(huffTable))
	binary.Write(file, binary.LittleEndian, numChars)

	for char, huffCode := range huffTable {
		binary.Write(file, binary.LittleEndian, uint32(char))
		binary.Write(file, binary.LittleEndian, uint8(huffCode.Bits))
		binary.Write(file, binary.LittleEndian, huffCode.Code) // Now writes uint64
		binary.Write(file, binary.LittleEndian, uint32(huffCode.Freq))
	}

	// Write the text length
	textLength := uint64(len([]rune(text)))
	binary.Write(file, binary.LittleEndian, textLength)

	// Write the compressed data
	file.Write(compressedData)
}

func CompressFile(text string, filename string) {
	frequencies := map[rune]int{}
	for _, ch := range text {
		frequencies[ch]++
	}

	fmt.Printf("Building huffman tree...\n")
	huffmanTree := buildHuffmanTree(frequencies)
	fmt.Println("Huffman tree built!")

	fmt.Printf("Building huffman table...\n")
	huffmanTable := buildHuffmanTable(huffmanTree)
	fmt.Println("Huffman table built!")

	fmt.Printf("Encoding data...\n")
	encodedData := encodeData(text, huffmanTable)
	fmt.Println("Data encoded!")

	fmt.Printf("Writing encoded data to file...\n")
	writeEncodedDataToFile(filename, text, huffmanTable, encodedData)
	fmt.Println("Encoded data written to file!")
}

/* Decompression and decoding functions */

func readEncodedDataFromFile(filename string) (map[rune]HuffCode, []byte, uint64) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the encoded Huffman table header
	var numChars uint32
	binary.Read(file, binary.LittleEndian, &numChars)

	huffTable := make(map[rune]HuffCode)
	for range numChars {
		var char uint32
		var bits uint8
		var code uint64 // Changed from byte to uint64
		var freq uint32
		binary.Read(file, binary.LittleEndian, &char)
		binary.Read(file, binary.LittleEndian, &bits)
		binary.Read(file, binary.LittleEndian, &code)
		binary.Read(file, binary.LittleEndian, &freq)

		huffTable[rune(char)] = HuffCode{Bits: int(bits), Code: code, Freq: int(freq)}
	}

	// Read the text length
	var textLength uint64
	binary.Read(file, binary.LittleEndian, &textLength)

	// Read the encoded data
	encodedData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return huffTable, encodedData, textLength
}

func buildDecodingTree(huffTable map[rune]HuffCode) *HuffTree {
	// Start with empty root
	root := &HuffIntermediateNode{weight: 0}

	// For each character, build its path in the tree
	for char, huffCode := range huffTable {
		currentNode := HuffBaseNode(root)

		// Traverse bits from MSB to LSB
		for i := huffCode.Bits - 1; i >= 0; i-- {
			bit := (huffCode.Code >> i) & 1

			// Check if currentNode is a leaf (shouldn't happen in valid Huffman tree)
			if currentNode.IsLeaf() {
				panic(fmt.Sprintf("Prefix conflict: trying to traverse through leaf for char %q (code: %08b, bits: %d)", char, huffCode.Code, huffCode.Bits))
			}

			intermediate := currentNode.(*HuffIntermediateNode)

			if i == 0 {
				// Last bit - create leaf node
				if bit == 0 {
					intermediate.left = &HuffLeafNode{
						element: char,
						weight:  huffCode.Freq,
					}
				} else {
					intermediate.right = &HuffLeafNode{
						element: char,
						weight:  huffCode.Freq,
					}
				}
			} else {
				// Not last bit - create/traverse intermediate nodes
				if bit == 0 {
					// Go left
					if intermediate.left == nil {
						intermediate.left = &HuffIntermediateNode{weight: 0}
					}
					currentNode = intermediate.left
				} else {
					// Go right
					if intermediate.right == nil {
						intermediate.right = &HuffIntermediateNode{weight: 0}
					}
					currentNode = intermediate.right
				}
			}
		}
	}

	return &HuffTree{root: root}
}

func decodeData(tree *HuffTree, encodedData []byte, textLength uint64) string {

	var result []rune
	currentNode := tree.Root()

	for _, b := range encodedData {
		for bitPos := 7; bitPos >= 0; bitPos-- {
			if uint64(len(result)) == textLength {
				return string(result)
			}

			bit := (b >> bitPos) & 1

			if !currentNode.IsLeaf() {
				intermediate := currentNode.(*HuffIntermediateNode)
				if bit == 0 {
					currentNode = intermediate.Left()
				} else {
					currentNode = intermediate.Right()
				}
			}

			if currentNode.IsLeaf() {
				leaf := currentNode.(*HuffLeafNode)
				result = append(result, leaf.Element())
				currentNode = tree.Root()
			}
		}
	}

	return string(result)
}

func writeDecodedDataToFile(filename string, decodedData string) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(decodedData)

	return err
}

func DecompressFile(encodedFilename string, decodedFilename string) {
	fmt.Printf("Read encoded data from file...\n")
	huffmanTable, encodedData, textLength := readEncodedDataFromFile(encodedFilename)
	fmt.Println("Encoded data read from file!")

	frequencies := map[rune]int{}
	for char, code := range huffmanTable {
		frequencies[char] = code.Freq
	}

	fmt.Printf("Build decoding tree from huffman table...\n")
	huffmanTree := buildDecodingTree(huffmanTable)
	fmt.Println("Decoding tree built from huffman table!")

	fmt.Printf("Decode data...\n")
	decodedData := decodeData(huffmanTree, encodedData, textLength)
	fmt.Println("Data decoded!")

	fmt.Printf("Writing decoded data to file...\n")
	writeDecodedDataToFile(decodedFilename, decodedData)
	fmt.Println("Decoded data written to file!")
}
