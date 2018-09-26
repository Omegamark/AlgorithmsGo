package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Find(value int) *Node {
	node := l.First()
	for {
		if node.next == nil {
			fmt.Println("Value not found.")
			break
		} else if node.value == value {
			println("Value found.")
			return node
		} else {
			node = node.Next()
		}
	}
	return nil
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	// n := l.First()
	// println(n.value)
	// for {
	// 	println(n.value)
	// 	n = n.Next()
	// 	if n == nil {
	// 		break
	// 	}
	// }
	// n = l.Last()
	// for {
	// 	println(n.value)
	// 	n = n.Prev()
	// 	if n == nil {
	// 		break
	// 	}
	// }
	// println("what am I?", n)
	l.Find(1)
}
