package main

import (
	"fmt"
	"hackerank/list"
)

func main() {
	l := list.List{}

	l.Append(1)
	l.Append(2)
	l.Append(42)
	l.Print()
	l.Prepend(0)
	l.Print()
	fmt.Println("Len:", l.Length())
	fmt.Println("100", l.Search(100))
	l.DeleteLast()
	l.Print()
	l.Append(6)
	l.Print()
	l.Delete(1)
	l.Print()
	l.DeleteFirst()
	l.Print()
}
