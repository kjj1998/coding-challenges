package huffman

type HuffBaseNode interface {
	IsLeaf() bool
	Weight() int
}

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

type HuffInternalNode struct {
	weight int
	left   HuffBaseNode
	right  HuffBaseNode
}

func (h *HuffInternalNode) Left() HuffBaseNode {
	return h.left
}

func (h *HuffInternalNode) Right() HuffBaseNode {
	return h.right
}

func (h *HuffInternalNode) Weight() int {
	return h.weight
}

func (h *HuffInternalNode) IsLeaf() bool {
	return false
}

type HuffTree struct {
	root HuffBaseNode
}

func (h *HuffTree) Root() HuffBaseNode {
	return h.root
}

func (h *HuffTree) Weight() int {
	return h.root.Weight()
}

func (h *HuffTree) PreOrderTraversal(visit func(HuffBaseNode, byte), code byte) {
	preOrderHelper(h.root, code, visit)
}

func preOrderHelper(node HuffBaseNode, code byte, visit func(HuffBaseNode, byte)) {
	if node == nil {
		return
	}

	visit(node, code)

	if !node.IsLeaf() {
		internal := node.(*HuffInternalNode)
		preOrderHelper(internal.left, (code << 1), visit)
		preOrderHelper(internal.right, (code<<1)|1, visit)
	}
}

type HuffCode struct {
	freq int
	code byte
	bits int
}
