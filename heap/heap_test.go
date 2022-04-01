package heap

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u *User) Value() float64 {
	return float64(u.Age)
}

func (u *User) Key() any {
	return u.ID
}

func TestInterface(t *testing.T) {
	heap := NewBinaryHeap(MaxHeap)
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

func (val value) Key() any {
	return float64(val)
}

func TestMaxHeapInsert(t *testing.T) {
	heap := NewBinaryHeap(MaxHeap)
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
	heap := NewBinaryHeap(MinHeap)
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

	heap := NewBinaryHeap(MaxHeap)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(l []value, heap *BinaryHeap) {
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

func TestMultiprocessMinHeapInsert(t *testing.T) {
	testdata := [][]value{
		{14, 45, 76, 33, 7, 10, 30, 9, 75, 50, 1},
		{56, 24, 77, 11, 9, 20, 90, 78, 71, 62},
		{12, 33, 4, 2, 0, 50, 60, 79, 65, 68},
	}
	result := []float64{0, 1, 2, 4, 7, 9, 9, 10, 11, 12, 14, 20, 24, 30, 33, 33, 45, 50, 50, 56, 60, 62, 65, 68, 71, 75, 76, 77, 78, 79, 90}

	heap := NewBinaryHeap(MinHeap)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(l []value, heap *BinaryHeap) {
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
	heap := NewBinaryHeap(MinHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	// result := []int{9, 30, 38, 43, 49, 55, 66, 72, 79, 83, 87, 91}
	for _, val := range testdata {
		heap.Insert(val)
	}

	v, err := heap.PopByIndex(5)
	if err != nil {
		t.Fatal(err)
	}

	heap2 := NewBinaryHeap(MinHeap)
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
	heap := NewBinaryHeap(MaxHeap)
	testdata := []value{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	for _, val := range testdata {
		heap.Insert(val)
	}

	v, err := heap.PopByIndex(5)
	if err != nil {
		t.Fatal(err)
	}

	heap2 := NewBinaryHeap(MaxHeap)
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
	heap := NewBinaryHeap(MaxHeap)
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
	heap := NewBinaryHeap(MinHeap)
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

type TValue struct {
	key int64
	val float64
}

func (tv *TValue) Key() interface{} {
	return tv.key
}

func (tv *TValue) Value() float64 {
	return tv.val
}

func TestFibHeapInsertPop(t *testing.T) {
	heap := NewFibHeap(MinHeap)
	rand.Seed(time.Now().UnixNano())
	var id int64
	var count int64
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				err := heap.Insert(&TValue{
					key: atomic.AddInt64(&id, 1),
					val: float64(rand.Intn(10000)),
				})
				if err != nil {
					log.Fatal(err)
				}
				if i%100 == 1 {
					heap.Pop()
					atomic.AddInt64(&count, 1)
				}
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)
	for {
		_, err := heap.Pop()
		if err != nil {
			break
		}
		count++
		// t.Log(n.Key(), n.Value())
	}

	t.Log(count)
}

func TestFibHeapUpdateValue(t *testing.T) {
	heap := NewFibHeap(MaxHeap)
	heap.Insert(&User{ID: 1, Name: "zhangsan", Age: 20})
	heap.Insert(&User{ID: 2, Name: "lisi", Age: 25})
	heap.Insert(&User{ID: 3, Name: "wangwu", Age: 22})
	heap.Insert(&User{ID: 4, Name: "zhaoliu", Age: 26})
	heap.Insert(&User{ID: 5, Name: "liuyi", Age: 26})
	heap.Insert(&User{ID: 6, Name: "chener", Age: 19})
	heap.Insert(&User{ID: 7, Name: "sunqi", Age: 30})
	heap.Insert(&User{ID: 8, Name: "zhouba", Age: 18})
	heap.Insert(&User{ID: 9, Name: "wujiu", Age: 29})
	heap.Insert(&User{ID: 10, Name: "zhengshi", Age: 35})

	u, err := heap.Pop()
	if err != nil || (u.Value() != 18 && heap.T() == MinHeap) || (u.Value() != 35 && heap.T() == MaxHeap) {
		t.Fail()
	}

	heap.UpdateValue(&User{ID: 4, Name: "zhaoliu", Age: 28})
	heap.UpdateValue(&User{ID: 7, Name: "sunqi", Age: 12})

	for val, err := heap.Pop(); err == nil; val, err = heap.Pop() {
		t.Log(val.(*User))
	}
}

func TestFibHeapDelete(t *testing.T) {
	heap := NewFibHeap(MinHeap)
	heap.Insert(&User{ID: 1, Name: "zhangsan", Age: 20})
	heap.Insert(&User{ID: 2, Name: "lisi", Age: 25})
	heap.Insert(&User{ID: 3, Name: "wangwu", Age: 22})
	heap.Insert(&User{ID: 4, Name: "zhaoliu", Age: 26})
	heap.Insert(&User{ID: 5, Name: "liuyi", Age: 26})
	heap.Insert(&User{ID: 6, Name: "chener", Age: 19})
	heap.Insert(&User{ID: 7, Name: "sunqi", Age: 30})
	heap.Insert(&User{ID: 8, Name: "zhouba", Age: 18})
	heap.Insert(&User{ID: 9, Name: "wujiu", Age: 29})
	heap.Insert(&User{ID: 10, Name: "zhengshi", Age: 35})

	u, err := heap.Pop()
	if err != nil || (u.Value() != 18 && heap.T() == MinHeap) || (u.Value() != 35 && heap.T() == MaxHeap) {
		t.Fail()
	}

	v := heap.Delete(4)
	t.Log("delete: ", v.(*User))
	for val, err := heap.Pop(); err == nil; val, err = heap.Pop() {
		if val.Key() == 4 {
			t.Fail()
		}
		t.Log(val.(*User))
	}
}
