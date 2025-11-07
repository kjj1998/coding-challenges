package huffman

import (
	"container/heap"
)

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

		internal := &HuffInternalNode{
			left:   left,
			right:  right,
			weight: left.Weight() + right.Weight(),
		}

		heap.Push(&pq, internal)
	}

	root := heap.Pop(&pq).(HuffBaseNode)

	return &HuffTree{root: root}
}
