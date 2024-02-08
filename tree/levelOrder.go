package main

import (
	"fmt"
)

type node struct {
	key   int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func (n *node) Insert(k int) {
	if n.key == 0 {
		n.key = k
	}

	if n.key < k {
		// Move right
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			n.right.Insert(k)
		}

	} else if n.key > k {
		// Move left
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// Pre Order Traversal
func (n *node) PreOrderTraversal() {
	if n != nil {
		fmt.Println(n.key)
		n.left.PreOrderTraversal()
		n.right.PreOrderTraversal()
	}
	return
}

func (root *node) levelOrder() {
	// q := deque.NewDeque[int]()
	// for root != nil {
	// node := q.PopFront()
	// if node != 0 {
	// 	fmt.Print(root.key)
	// 	q.PushBack(root.left.key)
	// 	q.PushBack(root.right.key)
	// }
	// }
}

// 1. bikin func untuk merge result yg sama kedalam queue
// 2. pop bagian depan queue(FIFO)
// 3. print result
func (root *node) tca(v int) {
	result := 0

	if root != nil {
		temp := root
		if root.left != nil || root.right != nil {
			if temp.key < v {
				result = temp.key
				temp.right.tca(v)
			} else {
				result = temp.key
				temp.left.tca(v)
			}
		}
		if result != 0 {
			fmt.Println(result)
		}
	}
}

func (root *node) lca(v1, v2 int) int {
	if v1 == root.key || v2 == root.key {
		return root.key
	}
	for root != nil {
		if v1 > root.key && v2 > root.key {
			root = root.right
		} else if v1 < root.key && v2 < root.key {
			root = root.left
		} else {
			return root.key
		}
	}
	return 0
}

func main() {
	// var q deque.Deque[string]
	// q.PushBack("foo")
	// q.PushBack("bar")
	// q.PushBack("baz")

	// fmt.Println(q.PopFront())
	tree := &node{}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(400)
	tree.Insert(500)
	tree.Insert(350)
	tree.Insert(20)
	tree.Insert(75)
	tree.Insert(10)
	tree.Insert(148)
	tree.Insert(25)
	// tree.levelOrder()
	// tree.tca(10)
	fmt.Println(tree.lca(148, 25))
	// tree.PreOrderTraversal()
}
