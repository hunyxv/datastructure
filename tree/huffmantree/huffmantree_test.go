package huffmantree

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"testing"
)

func print01(code int64) {
	bytes := []int{}
	for code != 0 {
		bit := int(code % 2)
		code >>= 1
		bytes = append(bytes, bit)
	}
	// -2 是因为要去掉开头占位的 “1”
	for i := len(bytes) - 2; i >= 0; i-- {
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

// 测试压缩一张 BMP 图片
type FileHeader struct {
	Head                 uint16
	ImageSize            uint32
	ReservedA, ReservedB uint16
	Offbits              uint32
}

type DIBHeader struct {
	Size        uint32
	Width       uint32
	Height      uint32
	PlaneNum    uint16
	PixelSize   uint16
	A_          uint32
	ImageSize   uint32
	XResolution uint32
	Resolution  uint32
	ColorNum    uint32
}

func statisticalFrequency(file io.Reader, pixelSize int) (freq map[uint32]float64) {
	freq = make(map[uint32]float64)
	g := pixelSize / 8
	pixel := make([]byte, g)
	var c uint32
	for {
		n, err := file.Read(pixel)
		if err == io.EOF {
			break
		}
		if n < g {
			panic(fmt.Errorf("err: %d", n))
		}
		for i := 0; i < g; i++ {
			c <<= 8
			c += uint32(pixel[i])
		}
		freq[c]++
	}

	fmt.Println(freq)
	for k, v := range freq {
		freq[k] = v / float64(len(freq))
	}
	return
}

// 这里只压缩像素数据，没有压缩文件头
func TestBMP(t *testing.T) {
	file, err := os.Open("bmp2.bmp")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	fileHeader := new(FileHeader)
	dibHeader := new(DIBHeader)
	binary.Read(file, binary.LittleEndian, fileHeader)
	binary.Read(file, binary.LittleEndian, dibHeader)
	file.Seek(int64(fileHeader.Offbits), 0)

	freq := statisticalFrequency(file, int(dibHeader.PixelSize))
	fmt.Println(len(freq), freq)

	l := make([]Element, 0)
	for k, v := range freq {
		e := &element{c: k, w: float64(v) / float64(len(freq))}
		t.Log(e.Char())
		l = append(l, e)
	}

	tree := CreateHuffmanTree(l...)
	tree.GenerateCode()
	var mask int64 = 255<<7
	var table = make(map[int32][]byte)
	tree.TraversalLeaf(func(leaf Leaf) {
		t.Logf("%d", leaf.Code())
		//table[int32(leaf.Code())] = 
		for i:=0;i<8;i++{

		}
	})

	compressedFile, err := os.OpenFile("bmp.huf", 066, os.FileMode(os.O_CREATE)|os.FileMode(os.O_WRONLY))
	if err != nil {
		t.Fatal(err)
	}
	defer compressedFile.Close()

	// 先写入文件头
	header := make([]byte, int(fileHeader.Offbits))
	file.Seek(0, 0)
	file.Read(header)
	compressedFile.Write(header)


}
