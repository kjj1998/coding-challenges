package huffman

/* Base Node */

type HuffBaseNode interface {
	IsLeaf() bool
	Weight() int
}

/* Leaf Node */

type HuffLeafNode struct {
	element rune
	weight  int
}

func (h *HuffLeafNode) Element() rune {
	return h.element
}

func (h *HuffLeafNode) IsLeaf() bool {
	return true
}

func (h *HuffLeafNode) Weight() int {
	return h.weight
}

/* Intermediate Node */

type HuffIntermediateNode struct {
	weight int
	left   HuffBaseNode
	right  HuffBaseNode
}

func (h *HuffIntermediateNode) Left() HuffBaseNode {
	return h.left
}

func (h *HuffIntermediateNode) Right() HuffBaseNode {
	return h.right
}

func (h *HuffIntermediateNode) Weight() int {
	return h.weight
}

func (h *HuffIntermediateNode) IsLeaf() bool {
	return false
}

/* Huffman Tree */

type HuffTree struct {
	root HuffBaseNode
}

func (h *HuffTree) Root() HuffBaseNode {
	return h.root
}

func (h *HuffTree) Weight() int {
	return h.root.Weight()
}

func (h *HuffTree) PreOrderTraversal(visit func(HuffBaseNode, byte, int, map[rune]HuffCode)) map[rune]HuffCode {
	var code byte = 0
	var bits int = 0
	huffTable := map[rune]HuffCode{}

	preOrderHelper(h.root, code, bits, huffTable, visit)

	return huffTable
}

func preOrderHelper(node HuffBaseNode, code byte, bits int, huffTable map[rune]HuffCode, visit func(HuffBaseNode, byte, int, map[rune]HuffCode)) {
	if node == nil {
		return
	}

	visit(node, code, bits, huffTable)

	if !node.IsLeaf() {
		intermediate := node.(*HuffIntermediateNode)
		preOrderHelper(intermediate.left, (code << 1), bits+1, huffTable, visit)
		preOrderHelper(intermediate.right, (code<<1)|1, bits+1, huffTable, visit)
	}
}

/* Huffman Code */

type HuffCode struct {
	Freq int
	Code byte
	Bits int
}
