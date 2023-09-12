package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) pop() {
	if l.head == nil {
		return
	} else if l.head.next == nil {
		l.head = nil
		return
	}
	curr := l.head
	next := l.head.next
	for next.next != nil {
		curr = next
		next = next.next
	}

	curr.next = nil
}

func (l *List) printList() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main() {

	list := &List{}

	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	list.pop()
	list.pop()

	list.printList()
}
