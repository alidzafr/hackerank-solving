package main

import (
	"fmt"
)

type listnode struct {
	Data int
	Next *listnode
	Prev *listnode
}

type llist struct {
	Head *listnode
}

func (l *llist) add(data int) *listnode {
	node := &listnode{Data: data}
	if l.Head == nil {
		l.Head = node
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
	return l.Head
}

func (l llist) Print() {
	p := l.Head
	fmt.Print("Head")
	for p != nil {
		fmt.Print("->", p.Data)
		p = p.Next
	}
	fmt.Printf("\n")
}

func arr_cycle(l *listnode) bool {

	tempArr := []int{}
	for l != nil {
		for i := range tempArr {
			if tempArr[i] == l.Data {
				return true
			}
		}
		tempArr = append(tempArr, l.Data)
		l = l.Next
	}
	return false
}

func has_cycle(l *listnode) bool {
	return false
}

func sortingInsert(l *listnode, data int) *listnode {
	node := &listnode{Data: data}
	temp := l

	if l.Data > data {
		l.Prev = node
		node.Next = l
		return node
	}

	for temp.Next != nil {
		if temp.Data < data {
			temp = temp.Next
		} else {
			temp.Prev.Next = node
			node.Prev = temp.Prev
			node.Next = temp
			temp.Prev = node
			return l
		}
	}

	if temp.Data < data {
		temp.Next = node
		node.Prev = temp
	} else {
		temp.Prev.Next = node
		node.Prev = temp.Prev
		temp.Prev = node
		node.Next = temp
	}
	return l
}

func main() {
	l := llist{}
	l.add(1)
	l.add(2)
	l.add(3)
	head := l.add(2)
	fmt.Println(has_cycle(head))
	// llist := l.add(4)
	// l.Print()
	// sortingInsert(llist, 3)
	l.Print()
}
