package list

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	Data int
	Next *SinglyLinkedListNode
}

type list struct {
	Head   *SinglyLinkedlistNode
	length int
}

func (l *list) Append(data int) {
	node := &SinglyLinkedListNode{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
}

func (l *list) InsertNodeAtPosition(data, position int) {
	curr := l.Head
	count := 1
	for count != position {
		curr = curr.Next
		count++
	}
	node := &SinglyLinkedListNode{Data: data, Next: nil}
	node.Next = curr.Next
	curr.Next = node
}

func (l *list) deleteNode(position int) {
	if l.Head == nil {
		return
	}
	curr := l.Head
	count := 0

	if position == 0 {
		l.Head = l.Head.Next
	}
	for curr.Next != nil {
		count++
		if count == position {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}

}

func (l list) Print() {
	p := l.Head
	fmt.Print("Head")
	for p != nil {
		if p.Data != 0 {
			fmt.Print("->", p.Data)
		}
		p = p.Next
	}
	fmt.Printf("\n")
}

func (l *list) Reverse() *SinglyLinkedListNode {
	prev := &SinglyLinkedListNode{}
	curr := l.Head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp

	}
	return prev

}

func main() {
	l := List{}

	var len, data int

	fmt.Print("lenght: ")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Scan(&data)
		l.Append(data)
	}
	l.Reverse()
	l.Print()
	// fmt.Scan(&data)
	// fmt.Scan(&position)
	// l.deleteNode(position)

}
