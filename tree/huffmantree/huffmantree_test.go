package huffmantree

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"syscall"
	"testing"
)

func print01(code uint64) {
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

// FileHeader 文件头 共 2 + 4 + 2 + 2 + 4 = 14 字节
type FileHeader struct {
	Head                 uint16 // 头文件字段， bmp 内容为 0x424D （ASCII 对应为 BM）
	ImageSize            uint32 // 整个 bmp 文件大小
	ReservedA, ReservedB uint16 // 预留字段 0
	Offbits              uint32 // 图片信息开始位置
}

// DIBHeader 共 40 字节 即 DIBHeader.Size
type DIBHeader struct {
	Size        uint32 // DIB header 大小
	Width       uint32 // 图片宽
	Height      uint32 // 图片高
	PlaneNum    uint16 // 色彩屏幕数量 1
	PixelSize   uint16 // 每个像素用用多少 bit 来表示
	A_          uint32 // 压缩方式，通常不压缩 BI_RGB 对应为0
	ImageSize   uint32 // 图片大小（原始位图数据的大小），对于不压缩的图片 通常为0
	XResolution uint32 // 横向分辨率 像素
	Resolution  uint32 // 纵向分辨率
	ColorNum    uint32 // 调色板中颜色数量，通常为0（0不表示没有颜色）
	ImpColorNum uint32 // 重要颜色数量，通常为0（表示每种颜色都重要）
}

// len(FileHeader) + len(DIBHeader) == 54 byte == FileHeader.Offbits

// 统计每种像素出现的频率
func statisticalFrequency(file io.Reader, pixelSize int) (freq map[uint32]float64) {
	freq = make(map[uint32]float64)
	g := pixelSize / 8
	pixel := make([]byte, g)

	for {
		var c uint32
		n, err := file.Read(pixel)
		if err == io.EOF {
			break
		}
		if n < g {
			panic(fmt.Errorf("err: %d", n))
		}
		for i := 0; i < g; i++ {
			c <<= 8
			c |= uint32(pixel[i])
		}
		freq[c]++
	}

	// fmt.Println("像素出现次数：", freq)
	var total float64
	for _, v := range freq {
		total += v
	}
	for k, v := range freq {
		freq[k] = v / total
	}
	// fmt.Println("像素出现频率：", freq)
	return
}

func uint64TO01Code(code uint64) byte {
	var b byte
	for code != 0 {
		bit := byte(code % 2)
		code >>= 1
		b |= bit
		b <<= 1
	}
	return b
}

func uint64TOChar(n uint32, pixelSize int) (b []byte) {
	b = make([]byte, 0, pixelSize)
	for i := 0; i < pixelSize/8; i++ {
		b = append([]byte{uint8(n & 255)}, b...)
		n >>= 8
	}
	return
}

// 这里只压缩像素数据，没有压缩文件头
func TestBMPCompress(t *testing.T) {
	// 源文件
	file, err := os.Open("testbmp/bmp.bmp")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	fileHeader := new(FileHeader)
	dibHeader := new(DIBHeader)
	binary.Read(file, binary.LittleEndian, fileHeader)
	binary.Read(file, binary.LittleEndian, dibHeader)
	file.Seek(int64(fileHeader.Offbits), 0)
	t.Logf("%+v\n%+v", fileHeader, dibHeader)

	// 计算每种颜色像素权重
	freq := statisticalFrequency(file, int(dibHeader.PixelSize))
	file.Seek(0, 0)

	// 赫夫曼树节点
	l := make([]Element, 0, len(freq))
	for k, v := range freq {
		e := &element{c: k, w: v}
		l = append(l, e)
	}
	// 创建赫夫曼树 并且生成赫夫曼树编码
	tree := CreateHuffmanTree(l...)
	tree.GenerateCode()

	var table = make(map[uint32]byte)
	tree.TraversalLeaf(func(leaf Leaf) {
		print01(leaf.Code())            // 打印赫夫曼树编码 （01显示）
		t.Log(leaf.Code(), leaf.Char()) // 打印赫夫曼树编码（uint64）和 像素（uint32表示）

		pixel := leaf.Char().(uint32)
		table[pixel] = uint64TO01Code(leaf.Code()) // 像素和赫夫曼编码对应关系
	})

	// 压缩文件
	compressedFile, err := os.OpenFile("testbmp/bmp_compress.huf", syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer compressedFile.Close()
	bufw := bufio.NewWriter(compressedFile)
	defer bufw.Flush()

	// 先写入文件头（头部未压缩）
	header := make([]byte, int(fileHeader.Offbits))
	_, err = file.Read(header)
	if err != nil {
		t.Fatal(err)
	}
	bufw.Write(header)

	bt, err := json.Marshal(table)
	if err != nil {
		t.Fatal(err)
	}
	// 写入 像素和赫夫曼树编码 对应关系
	if _, err := bufw.Write(bt); err != nil {
		t.Fatal(err)
	}
	bufw.WriteByte('\n')

	// 将每个像素替换为赫夫曼编码写入压缩文件
	g := int(dibHeader.PixelSize) / 8
	pixel := make([]byte, g)
	for {
		var c uint32
		n, err := file.Read(pixel)
		if err == io.EOF {
			break
		}
		if n != g {
			t.Fail()
		}
		for i := 0; i < g; i++ {
			c <<= 8
			c |= uint32(pixel[i])
		}
		bufw.WriteByte(table[c])
	}
}

// 解压缩
func TestBMPDecompress(t *testing.T) {
	// 压缩文件
	hufFile, err := os.Open("testbmp/bmp_compress.huf")
	if err != nil {
		t.Fatal(err)
	}
	hufBuf := bufio.NewReader(hufFile)

	// 解压缩文件
	bmpFile, err := os.OpenFile("testbmp/bmp_decompress.bmp", syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer hufFile.Close()
	defer bmpFile.Close()

	bmpBuf := bufio.NewWriter(bmpFile)
	defer bmpBuf.Flush()
	// 读取头部数据
	fileHeader := new(FileHeader)
	dibHeader := new(DIBHeader)
	binary.Read(hufFile, binary.LittleEndian, fileHeader)
	binary.Read(hufFile, binary.LittleEndian, dibHeader)
	hufFile.Seek(0, 0)

	// 读取bmp文件头 并写入解压缩文件
	header := make([]byte, int(fileHeader.Offbits))
	n, err := hufBuf.Read(header)
	if err != nil || n != int(fileHeader.Offbits) {
		t.Fail()
	}
	bmpBuf.Write(header)

	// 读取 【像素：赫夫曼树编码】 对应关系
	btable, err := hufBuf.ReadBytes('\n')
	if err != nil {
		t.Fatal(err)
	}

	tmp := make(map[uint32]byte)
	fmt.Println(string(btable))
	if err := json.Unmarshal(btable, &tmp); err != nil {
		t.Fatal(err)
	}

	// 【像素：赫夫曼树编码】 反转为 【赫夫曼树编码：像素】
	table := make(map[byte]uint32)
	for k, v := range tmp {
		table[v] = k
	}

	// 将赫夫曼编码替换为像素写入文件
	for {
		b, err := hufBuf.ReadByte()
		if err == io.EOF {
			break
		}

		bmpBuf.Write(uint64TOChar(table[b], int(dibHeader.PixelSize)))
	}
}
