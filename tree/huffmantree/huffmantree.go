package huffmantree

import (
	"github.com/hunyxv/datastructure/heap"
)

// Element .
type Element interface {
	Weight() float64
	Char() interface{}
}

// Node 普通节点
type Node interface {
	heap.Interface
	IsLeaf() bool
	Traversal(func(Node))
	TraversalLeaf(func(Leaf))
	GenerateCode()
	generateCode(uint64)
}

// Leaf 叶子节点
type Leaf interface {
	Node
	Char() interface{}
	Code() uint64
}

var _ Node = (*HuffmanNode)(nil)

// HuffmanNode 赫夫曼树节点
type HuffmanNode struct {
	w        float64
	isLeaf   bool
	lsubNode Node
	rsubNode Node
}

func newHuffmanNode(lnode, rnode Node) Node {
	return &HuffmanNode{
		w:        lnode.Value() + rnode.Value(),
		lsubNode: lnode,
		rsubNode: rnode,
	}
}

// Value 返回权重
func (n *HuffmanNode) Value() float64 {
	return n.w
}

func (n *HuffmanNode) Key() any { return nil }

// IsLeaf 是否为叶子节点
func (n *HuffmanNode) IsLeaf() bool {
	return n.isLeaf
}

// GenerateCode 遍历树生成赫夫曼编码
func (n *HuffmanNode) GenerateCode() {
	// 这个 1 是占位置的
	n.lsubNode.generateCode(1 << 1)
	n.rsubNode.generateCode(1<<1 + 1)
}

func (n *HuffmanNode) generateCode(code uint64) {
	n.lsubNode.generateCode(code << 1)
	n.rsubNode.generateCode(code<<1 + 1)
}

// Traversal 前序遍历各个节点
func (n *HuffmanNode) Traversal(f func(Node)) {
	f(n)
	n.lsubNode.Traversal(f)
	n.rsubNode.Traversal(f)
}

// TraversalLeaf 遍历赫夫曼树叶子节点，只在叶子节点执行 f
func (n *HuffmanNode) TraversalLeaf(f func(Leaf)) {
	n.lsubNode.TraversalLeaf(f)
	n.rsubNode.TraversalLeaf(f)
}

var _ Leaf = (*HuffmanLeafNode)(nil)

// HuffmanLeafNode 赫夫曼树叶子节点
type HuffmanLeafNode struct {
	w      float64
	char   interface{}
	code   uint64
	isLeaf bool
}

func newHuffmanLeafNode(weight float64, ch interface{}) *HuffmanLeafNode {
	return &HuffmanLeafNode{
		w:      weight,
		char:   ch,
		isLeaf: true,
	}
}

// Value 返回权重
func (n *HuffmanLeafNode) Value() float64 {
	return n.w
}

func (n *HuffmanLeafNode) Key() any { return nil }

// IsLeaf 是否为叶子节点
func (n *HuffmanLeafNode) IsLeaf() bool {
	return n.isLeaf
}

// Char 返回该节点保存的字符
func (n *HuffmanLeafNode) Char() interface{} {
	return n.char
}

// Code 返回字符的赫夫曼编码
func (n *HuffmanLeafNode) Code() uint64 {
	return n.code
}

// GenerateCode .
func (n *HuffmanLeafNode) GenerateCode() {}

func (n *HuffmanLeafNode) generateCode(code uint64) {
	n.code = code
}

// Traversal .
func (n *HuffmanLeafNode) Traversal(f func(Node)) {
	f(n)
}

// TraversalLeaf .
func (n *HuffmanLeafNode) TraversalLeaf(f func(Leaf)) {
	f(n)
}

// CreateHuffmanTree 创建赫夫曼树
func CreateHuffmanTree(elements ...Element) Node {
	if len(elements) == 0 {
		return nil
	}

	var node Node
	// 最小堆
	eleHeap := heap.NewBinaryHeap(heap.MinHeap)
	for _, ele := range elements {
		node = newHuffmanLeafNode(ele.Weight(), ele.Char())
		eleHeap.Insert(node)
	}

	// 构建赫夫曼树
	var leftTree, rightTree Node
	for {
		left, _ := eleHeap.Pop()
		leftTree = left.(Node)

		right, err := eleHeap.Pop()
		if err != nil {
			return leftTree
		}
		rightTree = right.(Node)
		newNode := newHuffmanNode(leftTree, rightTree)
		eleHeap.Insert(newNode)
	}
}
