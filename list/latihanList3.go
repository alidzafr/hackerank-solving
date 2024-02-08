package main

import "fmt"

type node struct {
	Data int
	Next *node
}

type list1 struct {
	Head *node
	Size int
}

func (l *list1) Add(data int) {
	dummyNode := &node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = dummyNode
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
			l.Size++
		}
		p.Next = dummyNode
	}
}

func (l *list1) getDataPosition(position int) int {
	temp := l.Head.Next
	count := 0

	for temp != nil {
		count++
		temp = temp.Next
	}

	for count > position {
		count -= 1
		l.Head = l.Head.Next
	}
	return l.Head.Data
}

func main() {
	l := list1{}

	var test, len, data, position int

	fmt.Scan(&test)
	for i := 0; i < test; i++ {
		fmt.Printf("\nlen:")
		fmt.Scan(&len)
		for i := 0; i < len; i++ {
			fmt.Scan(&data)
			l.Add(data)
		}
		fmt.Scan(&position)
		fmt.Println(l.getDataPosition(position))
	}
}
