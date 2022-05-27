package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

// 4
// 3

func (t *Tree) Insert(v int) {
	if t.Root == nil {
		t.Root = insert(t.Root, v)
		return
	}

	insert(t.Root, v)

}

func insert(n *Node, v int) *Node {
	if n == nil {
		newNode := &Node{
			Value: v,
			Left:  nil,
			Right: nil,
		}
		return newNode
	}

	if v < n.Value {
		n.Left = insert(n.Left, v)
	} else {
		n.Right = insert(n.Right, v)
	}

	return n
}

// main -> 0x00
// gua send -> Insert node 0x00
// m ->

type X struct {
}

func I(n *Node, v int) {
	if n == nil {
		n = &Node{
			Value: v,
			Left:  nil,
			Right: nil,
		}
		return
	}

	if v < n.Value {
		I(n.Left, v)
	} else {
		I(n.Right, v)
	}
}

func InOrder(n *Node) {
	if n == nil {
		return
	}

	InOrder(n.Left)
	fmt.Print(n.Value, " ")
	InOrder(n.Right)
}

//  tree -> 0x00
// root -> 0x01
func main() {
	var n int
	fmt.Scan(&n)

	tree := new(Tree)

	var d int
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		I(tree.Root, d)
	}

	fmt.Println("==== PRINT ====")
	InOrder(tree.Root)
}
