package main

import "fmt"

type linkedlistnode struct {
	Data int
	Next *linkedlistnode
}

type list1 struct {
	Head *linkedlistnode
}

type list2 struct {
	Head *linkedlistnode
}

func length(l *linkedlistnode) int {
	head := l
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}

func findMergePoint(A *linkedlistnode, B *linkedlistnode) int {

	tempArr := []int{}
	for A != nil {
		tempArr = append(tempArr, A.Data)
		A = A.Next
	}
	for B != nil {
		for i := 0; i < len(tempArr); i++ {
			if B.Data == tempArr[i] {
				return A.Data
			}
			B = B.Next
		}
	}
	return 0
}

// Function to reverse a given Linked List using Recursion
// func reverseList(head *linkedlistnode) *linkedlistnode {
// 	if head.Next == nil {
// 		return head
// 	}
// 	rest := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return rest
// }

func (l *list1) add(data int) *linkedlistnode {
	node := &linkedlistnode{Data: data, Next: nil}
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

func (l *list2) add(data int) *linkedlistnode {
	node := &linkedlistnode{Data: data, Next: nil}
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

func print(l *linkedlistnode) {
	fmt.Println("Head")
	for l != nil {
		fmt.Println(l.Data)
		l = l.Next
	}
	fmt.Printf("\n")
}

func sortedMerge(head1 *linkedlistnode, head2 *linkedlistnode) *linkedlistnode {
	result := &linkedlistnode{}
	curr_node := result

	for head1 != nil && head2 != nil {
		if head1.Data < head2.Data {
			curr_node.Next = head1
			head1 = head1.Next
		} else {
			curr_node.Next = head2
			head2 = head2.Next
		}
		curr_node = curr_node.Next
	}
	if head1 != nil {
		curr_node.Next = head1
		head1 = head1.Next
	} else if head2 != nil {
		curr_node.Next = head2
		head2 = head2.Next
	}

	return result.Next
}

// func sortedMerge(a *linkedlistnode, b *linkedlistnode) *linkedlistnode {
// 	a = reverseList(a)
// 	b = reverseList(b)

// 	var head, temp *linkedlistnode
// 	head = nil

// 	for a != nil && b != nil {
// 		if a.Data >= b.Data {
// 			temp = a.Next
// 			a.Next = head
// 			head = a
// 			a = temp
// 		} else {
// 			temp = b.Next
// 			b.Next = head
// 			head = b
// 			b = temp
// 		}
// 	}
// 	for a != nil {
// 		temp = a.Next
// 		a.Next = head
// 		head = a
// 		a = temp
// 	}
// 	for b != nil {
// 		temp = b.Next
// 		b.Next = head
// 		head = b
// 		b = temp
// 	}
// 	return head
// }

// func printList(node *linkedlistnode) {
// 	for node != nil {
// 		node.data
// 		node = node.Next
// 	}
// }

func main() {
	a := list1{}
	b := list2{}
	var test, data, len int

	fmt.Scan(&test)
	for x := 0; x < test; x++ {
		fmt.Println("len:")
		fmt.Scan(&len)
		for i := 0; i < len; i++ {
			fmt.Scan(&data)
			a.add(data)
		}
		print(a.Head)

		fmt.Println("len:")
		fmt.Scan(&len)
		for i := 0; i < len; i++ {
			fmt.Scan(&data)
			b.add(data)
		}
		print(b.Head)
	}
	list3 := sortedMerge(a.Head, b.Head)
	print(list3)
	// fmt.Scan(&list1)
	// fmt.Scan(&list2)

	// list3 = sortedMerge(a, b)

}
