package tree

import "fmt"

var count int

// Node represent the component of a binary search tree
type node struct {
	key   int
	left  *node
	right *node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
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

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.key < k {
		// move right
		return n.right.Search(k)
	} else if n.key > k {
		// Move left
		return n.left.Search(k)
	}
	return true
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

// Post Order Traversal
func (n *node) PostOrderTraversal() {
	if n != nil {
		n.left.PreOrderTraversal()
		n.right.PreOrderTraversal()
		fmt.Println(n.key)
	}
	return
}

// In Order Traversal
func (n *node) InOrderTraversal() {
	if n != nil {
		n.left.PreOrderTraversal()
		fmt.Println(n.key)
		n.right.PreOrderTraversal()
	}
	return
}

// Print right side top nodes
func (root *BinaryTree) printRightSide(node *TreeNode, depth int) {
	if node != nil {
		if root.level < depth {
			// Print the resultant node
			fmt.Print("  ", node.data)
			root.level = depth
		}
		// Recursively
		// First visit right subtree
		root.printRightSide(node.right, depth+1)
		// Then
		// Visit left subtree
		root.printRightSide(node.left, depth-1)
	}
}

// this is handle the request of printing top view
func (root *BinaryTree) topView() {
	root.level = -1
	root.printleftSide(root.root, 0)
	root.level = 0
	root.printRightSide(root.root, 0)
}

func main() {
	tree := &node{}
	// fmt.Println(tree)
	// tree.Insert(200)
	// tree.Insert(300)
	// tree.Insert(400)
	// tree.Insert(20)
	// tree.Insert(75)
	// tree.Insert(148)

	// fmt.Println(tree.Search(148))
	// fmt.Println(count)

	// tree.PreOrderTraversal()

	var len, key int

	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Scan(&key)
		tree.Insert(key)
	}
	tree.PostOrderTraversal()

}
