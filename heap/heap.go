package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type maxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *maxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *maxHeap) extract() int {
	extracted := h.array[0]

	l := len(h.array) - 1
	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will hapify from bottm top
func (h *maxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index) //update index agar loop berhenti
	}
}

// maxHeapifyDown will hapify top to bottom
func (h *maxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//  loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // When left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return i*2 + 1
}

// get the right child index
func right(i int) int {
	return i*2 + 2
}

// swap key in the array
func (h *maxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// delete element of index
func (h *maxHeap) delete(e int) {
	for i, element := range h.array {
		if element == e {
			h.array = append(h.array[:i], h.array[i+1:]...)
		}
	}
}

// check array has contains element or not
// so it act like .set()
func (h *maxHeap) has(element int) bool {
	// _, ok := h.array[element]
	// return ok
	for _, e := range h.array {
		if e == element {
			return true
		}
	}
	return false
}

// func (h *maxHeap) len() int {

// }

// NOT WORKING !!
func median() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	g := &maxHeap{}
	s := &maxHeap{}

	result := []float32{}

	for i := range a {
		g.insert(a[i])
		// x := g.extract()
		s.insert(g.extract())

		if len(g.array) > len(s.array) {
			// x = heap.Pop(g)
			s.insert(g.extract())
		}
		if len(g.array) != len(s.array) {
			// y := (*s)[0]
			result = append(result, float32(s.array[0]))
		} else {
			// result = append(result, float32(((*g)[0]-(*s)[0])/2))
			result = append(result, float32(((g.array)[0]-(s.array)[0])/2))
		}
	}
	for i := range result {
		fmt.Printf("%v\n", result[i])
	}
}

// NOT WORKING !!!
func printValue() {
	m := &maxHeap{}
	lookup := []int{}
	// fmt.Println(m)

	// buildHeap := []int{10, 29, 39, 5, 7, 9, 11, 13, 15, 17}
	// for _, v := range buildHeap {
	// 	m.insert(v)
	// 	fmt.Println(m)
	// }

	// for i := 0; i < 5; i++ {
	// 	m.extract()
	// 	fmt.Println(m)
	// }

	scanner := bufio.NewScanner(os.Stdin)
	operator := 0
	fmt.Print("number of query:")
	fmt.Scan(&operator)

	for i := 0; i < operator; i++ {
		scanner.Scan()
		s := scanner.Text()

		// make sure add format was correct
		if strings.Contains(s, "1") || strings.Contains(s, "2") || strings.Contains(s, "4") {
			for len(s) < 3 {
				fmt.Println("please input correct value")
				scanner.Scan()
				s = scanner.Text() //update string operator
				if strings.Contains(s, "3") {
					return
				}
			}
		}

		// Split operator and input [o ii]
		input := strings.Split(s, " ")

		// add element
		if input[0] == "1" {
			v, err := strconv.Atoi(input[1])
			if err != nil {
				panic(err)
			}
			set := m.has(v)
			if set == false {
				m.insert(v)
				lookup = append(lookup, v)
			}
			fmt.Println(m)

			// for _, e := range lookup {
			// 	if e == v {
			// 		break
			// 	} else {
			// 		m.insert(v)
			// 		lookup = append(lookup, v)
			// 	}
			// }
			// fmt.Println(m)

			// delete element
		} else if input[0] == "2" {
			v, err := strconv.Atoi(input[1])
			if err != nil {
				panic(err)
			}
			// m.delete(v)
			// fmt.Println(m)

			for i, e := range lookup {
				if e == v {
					lookup = append(lookup[:i], lookup[i+1:]...)
				}
			}

			// print lowest element
		} else if input[0] == "3" {
			// heapArr := m.array
			// fmt.Println("lowest value:", heapArr[len(heapArr)-1])

			for i, e := range m.array {
				for _, x := range lookup {
					if e != x {
						m.array = append(m.array[:i], m.array[i+1:]...)
					}
				}
			}
			fmt.Println(m.array[len(m.array)-1])

			// check element if already exist before inserting
		} else if input[0] == "4" {
			v, err := strconv.Atoi(input[1])
			if err != nil {
				panic(err)
			}
			ans := m.has(v)
			fmt.Println(ans)
		}

	}

}
