package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e:= l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	ls := list.New()
	s4 := ls.PushBack('d')
	s1 := ls.PushFront('a')
	ls.InsertBefore('c', s4)
	ls.InsertAfter('b', s1)
	for s := ls.Front(); s != nil; s = s.Next() {
		fmt.Printf("%c\n", s.Value)
	}
}
