package main

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) *Tree {
	// On first insert, if a Node is not extant, then create the first node.
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		// On subsequent inserts, simply insert the value at the right or left node.
		t.node.insert(value)
	}

	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			// Since left is a node, we can recursively use the node insert method.
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			// Since right is a node, we can recursively use the node insert method.
			n.right = &Node{value: value}
		} else {
			// This continues to call the insert method until it reaches a node with a value of nil, at which point it creates a new node with the inserted value.
			n.right.insert(value)
		}
	}
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}

	// Checking that the function is not traversing the entire tree.
	println(n.value)

	if value < n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}

	// println(n.value)
	printNode(n.left)
	printNode(n.right)

}

func main() {
	t := &Tree{}
	// We can daisy chain the insert method here since the t.insert method returns a *Tree which also has an insert method.
	t.insert(10).insert(8).insert(20).insert(9).insert(0).insert(25)

	printNode(t.node)
	println(t.node.exists(55))

}
