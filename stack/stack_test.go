package stack

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	stack := NewStack(10)
	for i := 0; i < 10; i++ {
		if err := stack.Push(i); err != nil {
			t.Fatal(err)
		}
	}
	for {
		_, err := stack.Pop()
		if err != nil {
			break
		}
	}
	if stack.StackLength() != 0 {
		t.Fatal("stack length should be 0")
	}
}

func TestStackTraverse(t *testing.T) {
	stack := NewStack(10)
	for i := 0; i < 10; i++ {
		if err := stack.Push(i); err != nil {
			t.Fatal(err)
		}
	}

	stack.StackTraverse(func(el interface{}) bool {
		i := el.(int)
		if i > 7 {
			return false
		}
		return true
	})
}

// 使用 stack 来迷宫求解
// 迷宫范围 1,1 --> 8,8
var maze = [10][10]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 0, 0, 1, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
	{1, 1, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

type point struct {
	x, y, c int
}

func (p *point) next() *point {
	p.c++
	var next *point
	switch p.c {
	case 1:
		next = &point{x: p.x + 1, y: p.y}
	case 2:
		next = &point{x: p.x, y: p.y - 1}
	case 3:
		next = &point{x: p.x - 1, y: p.y}
	case 4:
		next = &point{x: p.x, y: p.y + 1}
	default:
		return nil
	}
	if maze[next.x][next.y] == 0 {
		return next
	}
	return p.next()
}

func TestMaze(t *testing.T) {
	stack := NewStack(64)
	var count int
	p := &point{x: 1, y: 1}

	var _tk = true
	for _tk || !stack.IsEmpty() {
		_tk = false
		if maze[p.x][p.y] == 0 {
			stack.Push(p)
			maze[p.x][p.y] = 2
		}
		if p.x == 8 && p.y == 8 {
			count++
			if top, err := stack.Pop(); err != nil {
				break
			} else {
				t := top.(*point)
				maze[t.x][t.y] = 0
			}
			if top, err := stack.GetTop(); err == nil {
				p = top.(*point)
			}
			continue
		}

		if next := p.next(); next != nil && maze[next.x][next.y] == 0 {
			p = next
		} else {
			if top, err := stack.Pop(); err != nil {
				break
			} else {
				t := top.(*point)
				maze[t.x][t.y] = 0
			}

			if top, err := stack.GetTop(); err == nil {
				p = top.(*point)
			}
		}
	}

	t.Logf("迷宫有 %d 条路径通向出口", count)
}
