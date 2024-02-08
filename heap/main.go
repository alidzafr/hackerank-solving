package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) has(element int) bool {
	for _, v := range *h {
		if v == element {
			return true
		}
	}
	return false
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func qHeap1() {
	h := &IntHeap{}
	// heap.Init(h)
	// heap.Push(h, 3)
	// heap.Push(h, 5)
	// heap.Push(h, 4)
	// heap.Push(h, 2)
	// heap.Push(h, 1)
	// heap.Push(h, 6)
	// heap.Push(h, -1)
	// fmt.Println((*h)[0])
	// check := h.has(6)
	// fmt.Println(check)
	// fmt.Printf("minimum: %d\n", (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(h))
	// }

	scanner := bufio.NewScanner(os.Stdin)
	query := 0

	// fmt.Print("insert query: ")
	fmt.Scan(&query)
	for i := 0; i < query; i++ {
		scanner.Scan()
		// Turn input operator into array
		lookup := strings.Fields(scanner.Text())

		// Insert into heap
		if lookup[0] == "1" {
			v, _ := strconv.Atoi(lookup[1])
			if h.has(v) == false {
				heap.Push(h, v)
			}

			// Delete element from heap
		} else if lookup[0] == "2" {
			v, _ := strconv.Atoi(lookup[1])
			if h.has(v) == true {
				// Find the index
				for index, element := range *h {
					if element == v {
						v = index
						// fmt.Println(index)
					}
				}
				heap.Remove(h, v)
			}
			// Print
		} else if lookup[0] == "3" {
			fmt.Println((*h)[0])
		}
		fmt.Println(*h)
	}

}

func main() {
	h := &IntHeap{}
	// scanner := bufio.NewScanner(os.Stdin)

	// // Insert size and limit value (k)
	// scanner.Scan()
	// lookup := strings.Fields(scanner.Text())

	// query, _ := strconv.Atoi(lookup[0])
	// k, _ := strconv.Atoi(lookup[1])

	// // Insert query
	// element := 0
	// for i := 0; i < query; i++ {
	// 	fmt.Scan(&element)
	// 	heap.Push(h, element)
	// }

	k := 105823341
	// Insert query array
	A := []int{}
	for i := 0; i < 100000; i++ {
		A = append(A, 1)
	}

	for i := range A {
		heap.Push(h, A[i])
	}
	fmt.Println("heap after push: ", h)

	index := 0 // Jumlah percobaan
	head := (*h)[0]
	for i := 0; i < h.Len(); i++ {
		if head < int(k) {
			a := (*h)[0]
			heap.Pop(h)
			b := (*h)[0]
			heap.Pop(h)

			v := a + b*2
			heap.Push(h, v)
			i--

			head = (*h)[0]
			index++
		}
		if head == k && h.Len() == 1 {
			break
			// return index
		}
		if h.Len() < 2 {
			index = -1
		}
	}
	fmt.Println("result ", h)
	fmt.Println("index ", index)
}
