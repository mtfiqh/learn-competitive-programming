package main

import "fmt"

type Value int

type Node struct {
	Value Value
	Next  *Node
}

type List struct {
	First *Node
}

func (l *List) Insert(value Value) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}

	if l.First == nil {
		l.First = &newNode
		return
	}

	tail := FindTail(l.First)
	tail.Next = &newNode
	return
}

func (l *List) Print() {
	l.print(l.First)
}

func (l List) print(node *Node) {
	fmt.Println(node.Value)
	if node.Next != nil {
		l.print(node.Next)
	}
}

func FindTail(node *Node) *Node {
	if node.Next == nil {
		return node
	}

	return FindTail(node.Next)
}

func main() {
	var n int
	list := new(List)
	fmt.Scan(&n)

	var d int
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		list.Insert(Value(d))
	}

	list.Print()

}
