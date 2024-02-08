package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	a := l.PushBack(4)
	b := l.PushFront(1)
	l.InsertAfter(2, b)
	l.InsertBefore(3, a)

	// for i := l.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	p := l.Back()
	fmt.Print("Tail")
	for p != nil {
		fmt.Print("->", p.Value)
		p = p.Prev()
	}
	fmt.Println("")

}

DoublyLinkedListNode prev_node = null;
DoublyLinkedListNode curr = llist;

while(curr != null) {
	prev_node = curr.prev;
	curr.prev = curr.next;
	curr.next = prev_node;
	curr = curr.prev;
}
if(prev_node !=null){
	llist = prev_node.prev;
}
return llist;