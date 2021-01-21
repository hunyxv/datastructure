package heap

import (
	"sync"
	"testing"
)

type User struct {
	ID int 
	Name string
	Age int
}

func (u *User) Value() float64 {
	return float64(u.Age)
}

func TestInterface(t *testing.T) {
	heap := NewHeap(MaxHeap)
	heap.Insert(&User{ID: 1, Name: "zhangsan", Age: 20})
	heap.Insert(&User{ID: 2, Name: "lisi", Age: 25})
	heap.Insert(&User{ID: 3, Name: "wangwu", Age: 22})
	heap.Insert(&User{ID: 4, Name: "maliu", Age: 26})

	val, err := heap.Peek()
	if err != nil {
		t.Fatal(err)
	}

	user := val.(*User)
	if user.Name != "maliu" {
		t.FailNow()
	}
}

type value float64

func (val value) Value() float64 {
	return float64(val)
}
func TestMaxHeapInsert(t *testing.T) {
	heap := NewHeap(MaxHeap)
	result := []float64{91, 87, 83, 79, 72, 66, 55, 49, 43, 38, 30, 9}
	for _, val := range []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9} {
		heap.Insert(val)
	}

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}

func TestMinHeapInsert(t *testing.T) {
	heap := NewHeap(MinHeap)
	result := []float64{9, 30, 38, 43, 49, 55, 66, 72, 79, 83, 87, 91}
	for _, val := range []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9} {
		heap.Insert(val)
	}

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}

func TestMultiprocessMaxHeapInsert(t *testing.T) {
	testdata := [][]value{
		{14, 45, 76, 33, 7, 10, 30, 9, 75, 50, 1},
		{56, 24, 77, 11, 9, 20, 90, 78, 71, 62},
		{12, 33, 4, 2, 0, 50, 60, 79, 65, 68},
	}
	result := []float64{90, 79, 78, 77, 76, 75, 71, 68, 65, 62, 60, 56, 50, 50, 45, 33, 33, 30, 24, 20, 14, 12, 11, 10, 9, 9, 7, 4, 2, 1, 0}

	heap := NewHeap(MaxHeap)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(l []value, heap *Heap) {
			for _, val := range l {
				heap.Insert(val)
			}
			wg.Done()
		}(testdata[i], heap)
	}
	wg.Wait()

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}

func TestMultiprocessMInHeapInsert(t *testing.T) {
	testdata := [][]value{
		{14, 45, 76, 33, 7, 10, 30, 9, 75, 50, 1},
		{56, 24, 77, 11, 9, 20, 90, 78, 71, 62},
		{12, 33, 4, 2, 0, 50, 60, 79, 65, 68},
	}
	result := []float64{0, 1, 2, 4, 7, 9, 9, 10, 11, 12, 14, 20, 24, 30, 33, 33, 45, 50, 50, 56, 60, 62, 65, 68, 71, 75, 76, 77, 78, 79, 90}

	heap := NewHeap(MinHeap)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(l []value, heap *Heap) {
			for _, val := range l {
				heap.Insert(val)
			}
			wg.Done()
		}(testdata[i], heap)
	}
	wg.Wait()

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}

func TestMinHeapPopByIndex(t *testing.T) {
	heap := NewHeap(MinHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	// result := []int{9, 30, 38, 43, 49, 55, 66, 72, 79, 83, 87, 91}
	for _, val := range testdata {
		heap.Insert(val)
	}

	v, err := heap.PopByIndex(5)
	if err != nil {
		t.Fatal(err)
	}

	heap2 := NewHeap(MinHeap)
	for _, val := range testdata {
		if val.Value() != v.Value() {
			heap2.Insert(val)
		}
	}

	for {
		v1, err := heap.Pop()
		if err != nil {
			break
		}

		v2, err := heap2.Pop()
		if err != nil {
			break
		}

		if v1.Value() != v2.Value() {
			t.Fatalf("v1: %f, v2: %f", v1.Value(), v2.Value())
		}
	}
}

func TestMaxHeapPopByIndex(t *testing.T) {
	heap := NewHeap(MaxHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	for _, val := range testdata {
		heap.Insert(val)
	}

	v, err := heap.PopByIndex(5)
	if err != nil {
		t.Fatal(err)
	}

	heap2 := NewHeap(MaxHeap)
	for _, val := range testdata {
		if val.Value() != v.Value() {
			heap2.Insert(val)
		}
	}

	for {
		v1, err := heap.Pop()
		if err != nil {
			break
		}

		v2, err := heap2.Pop()
		if err != nil {
			break
		}

		if v1.Value() != v2.Value() {
			t.Fatalf("v1: %f, v2: %f", v1.Value(), v2.Value())
		}
	}
}

func TestMaxHeapReplace(t *testing.T) {
	heap := NewHeap(MaxHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	result := []float64{91, 87, 83, 79, 72, 66, 55, 49, 43, 30, 15, 9}

	for _, val := range testdata {
		heap.Insert(val)
	}

	err := heap.Replace(6, value(15))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}

func TestMinHeapReplace(t *testing.T) {
	heap := NewHeap(MinHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	result := []float64{9, 15, 30, 38, 43, 49, 55, 72, 79, 83, 87, 91}

	for _, val := range testdata {
		heap.Insert(val)
	}

	err := heap.Replace(6, value(15))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; ; i++ {
		val, err := heap.Pop()
		if err != nil {
			break
		}

		if val.Value() != result[i] {
			t.Fatalf("%f --> %f\n", result[i], val.Value())
		}
	}
}
