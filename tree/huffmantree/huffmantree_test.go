package huffmantree

import (
	"fmt"
	"testing"
)

func print01(code int64) {
	bytes := []int{}
	for code != 0 {
		bit := int(code % 2)
		code >>= 1
		bytes = append(bytes, bit)
	}
	for i := len(bytes) - 1; i >= 0; i-- {
		fmt.Print(bytes[i])
	}
	fmt.Println()
}

func TestPrint01(t *testing.T) {
	print01(10)
}

type element struct {
	c interface{}
	w float64
}

func (e *element) Char() interface{} {
	return e.c
}

func (e *element) Weight() float64 {
	return e.w
}

func TestCreateTree(t *testing.T) {
	data := "ABACCDA"
	count := make(map[rune]int)

	for _, r := range data {
		count[r]++
	}

	l := make([]Element, 0)
	for k, v := range count {
		e := &element{c: k, w: float64(v) / float64(len(count))}
		t.Log(e.Char())
		l = append(l, e)

	}

	tree := CreateHuffmanTree(l...)
	tree.GenerateCode()

	tree.TraversalLeaf(func(leaf Leaf) {
		print01(leaf.Code())
	})
}

// 测试压缩一张图片
