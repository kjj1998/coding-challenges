package huffman

import (
	"container/heap"
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
