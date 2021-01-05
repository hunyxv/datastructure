package lists

import (
	"fmt"
	"testing"
)

var Lists *GLNode

func printList(l Interface) {
	if l.GetTag() == Elem {
		fmt.Printf(" %v ", l.GetElem())
	} else {
		fmt.Printf("(")
		head, _ := l.GetHead()
		printList(head)
		fmt.Printf(")")
	}
	if next, has := l.Next(); has {
		printList(next)
	}
}

func TestGL1Add(t *testing.T) {
	lists := &GLNode{ElemTag: SubList}
	lists.Add(1)
	lists.Add("2")
	lists.Add([3]int{3, 3, 3})
	subLists := &GLNode{ElemTag: SubList}
	subLists.Add("sub_1")
	subLists.Add([]int{2, 3, 4, 5})
	lists.Add(subLists)
	printList(lists)
	Lists = lists
}

func TestGL1Traverse(t *testing.T) {
	Lists.Traverse(func(i interface{}) bool {
		t.Logf("%v\n", i)
		return true
	})
}

func TestGL1GetHead(t *testing.T) {
	head, err := Lists.GetHead()
	if err != nil {
		t.Fatal(err)
	}
	head, err = head.GetHead()
	if err != nil {
		t.Fatal(err)
	}
	printList(head)
}

func TestGl1GetTail(t *testing.T) {
	tail, err := Lists.GetTail()
	if err != nil {
		t.Fatal(err)
	}
	printList(tail)
}

func TestGL1Copy(t *testing.T) {
	list := Lists.CopyGList()
	printList(list)
	fmt.Println()
}

func TestGL1Delete(t *testing.T) {
	Lists.Delete(3)
	printList(Lists)
}

func TestGL1Insert(t *testing.T) {
	Lists.Insert(3, &GLNode{ElemTag: Elem, Atom: "insert"})
	printList(Lists)
	fmt.Println()
}

func TestGL1Length(t *testing.T) {
	l := Lists.GListLength()
	if l != 4 {
		t.Fatal(l)
	}
}

func TestGL1Depth(t *testing.T) {
	depth := Lists.GListDepth()
	if depth != 2 {
		t.Fatal(depth)
	}
	lists := &GLNode{ElemTag: SubList}
	lists.Add(Lists)
	lists.Add("abc")
	lists.Add(struct{ a int }{a: 10})
	depth = lists.GListDepth()
	if depth != 3 {
		t.Fatal(depth)
	}
}
