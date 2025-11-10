package huffman

import (
	"container/heap"
	"encoding/binary"
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

/* Huffman tree/table functions */

func BuildHuffmanTree(frequencies map[rune]int) *HuffTree {
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

func BuildHuffmanTable(huffTree *HuffTree) map[rune]HuffCode {
	huffTable := huffTree.PreOrderTraversal(func(node HuffBaseNode, code byte, bits int, huffTable map[rune]HuffCode) {
		if node.IsLeaf() {
			leaf := node.(*HuffLeafNode)
			huffCode := HuffCode{Freq: leaf.Weight(), Code: code, Bits: bits}
			huffTable[leaf.Element()] = huffCode
		}
	})

	return huffTable
}

func WriteEncodedFile(filename string, huffTable map[rune]HuffCode) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numChars := uint32(len(huffTable))
	binary.Write(file, binary.LittleEndian, numChars)

	for char, huffCode := range huffTable {
		binary.Write(file, binary.LittleEndian, uint32(char))
		binary.Write(file, binary.LittleEndian, uint8(huffCode.Bits))
		binary.Write(file, binary.LittleEndian, huffCode.Code)
		binary.Write(file, binary.LittleEndian, uint32(huffCode.Freq))
	}
}

func ReadEncodedFile(filename string) map[rune]HuffCode {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numChars uint32
	binary.Read(file, binary.LittleEndian, &numChars)

	huffTable := make(map[rune]HuffCode)
	for range numChars {
		var char uint32
		var bits uint8
		var code byte
		var freq uint32
		binary.Read(file, binary.LittleEndian, &char)
		binary.Read(file, binary.LittleEndian, &bits)
		binary.Read(file, binary.LittleEndian, &code)
		binary.Read(file, binary.LittleEndian, &freq)

		huffTable[rune(char)] = HuffCode{Bits: int(bits), Code: code, Freq: int(freq)}
	}

	return huffTable
}
