package main

import (
	"fmt"
)

type linkedNode struct {
	data int
	next *linkedNode
}

type llist1 struct {
	head *linkedNode
}

type llist2 struct {
	head *linkedNode
}

// type llist3 struct {
// 	head *linkedNode
// }

func (l *llist1) add(data int) *linkedNode {
	node := &linkedNode{data: data, next: nil}
	if l.head == nil {
		l.head = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
	return node
}

func (l *llist2) add(data int) *linkedNode {
	node := &linkedNode{data: data, next: nil}
	if l.head == nil {
		l.head = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
	return node
}

// func (l *llist3) add(data int) {
// 	node := &linkedNode{data: data, next: nil}
// 	if l.head == nil {
// 		l.head = node
// 	} else {
// 		p := l.head
// 		for p.next != nil {
// 			p = p.next
// 		}
// 		p.next = node
// 	}
// }

func compare(llist1 *linkedNode, llist2 *linkedNode) bool {
	for llist1 != nil || llist2 != nil {
		if llist1 == nil || llist2 == nil {
			return false
		}
		if llist1.data != llist2.data {
			return false
		} else {
			llist1 = llist1.next
			llist2 = llist2.next
		}
	}
	return true
}

// Gunakan logika pengulangan array 2 dimensi
func merge(head1 *linkedNode, head2 *linkedNode) *linkedNode {
	var temp, result, head3 *linkedNode

	for head1 != nil && head2 != nil {
		if head1.data <= head2.data {
			temp = head1
			head1 = head1.next
		} else {
			temp = head2
			head2 = head2.next
		}
		if result == nil {
			result = temp
			head3 = result
		} else {
			result.next = temp
			result = result.next
		}
		result.next = nil
	}
	if head1 != nil {
		result.next = head1
	} else {
		result.next = head2
	}
	return head3
}

func print(l *linkedNode) {
	p := l
	fmt.Print("Head")
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
	fmt.Printf("\n")
}

func main() {
	a := llist1{}
	b := llist2{}
	// c := llist3{}

	var test, len, data int
	var llist1, llist2 *linkedNode

	fmt.Println("Test:")
	fmt.Scan(&test)
	for x := 0; x < test; x++ {

		fmt.Println("Length:")
		fmt.Scan(&len)
		for i := 0; i < len; i++ {
			fmt.Scan(&data)
			a.add(data)
		}
		print(a.head)

		fmt.Println("Length:")
		fmt.Scan(&len)
		for i := 0; i < len; i++ {
			fmt.Scan(&data)
			b.add(data)
		}
		print(b.head)
		fmt.Println(compare(llist1, llist2))
		list3 := merge(a.head, b.head)
		print(list3)
		a.head = nil
		b.head = nil

	}
	// fmt.Println("print:", a.add(0, 0))
}
