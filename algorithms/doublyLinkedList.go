package main

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) add(value int) {
	newNode := &Node{data: value}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	curr := ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	newNode.prev = curr
}

func (ll *LinkedList) delLast() {
	if ll.head == nil {
		return
	}
	if ll.head.next == nil {
		ll.head = nil
	}

	curr := ll.head

	for curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil
}

func (ll *LinkedList) findPosition(value int) int {
	curr := ll.head

	for i := 0; curr != nil; i++ {
		if curr.data == value {
			return i
		}
		curr = curr.next
	}
	return -1

}

func (ll *LinkedList) printList() {
	curr := ll.head

	for curr != nil {
		println(curr.data)
		curr = curr.next
	}
}

func (ll *LinkedList) delfirstVal(value int) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.data == value {
		ll.head = ll.head.next
	}
	curr := ll.head.next

	for curr.next != nil {
		if curr.data == value {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			return true
		}
	}
	if curr.data == value {
		curr.prev = nil
		return true
	}

	return false
}

func main() {
	linkedList := &LinkedList{}

	linkedList.add(1)
	linkedList.add(2)
	linkedList.add(3)
	linkedList.add(4)

	linkedList.delLast()
	println("Print list")
	linkedList.printList()
	println("\nPrint find position of 2")
	println(linkedList.findPosition(2))

	println("\n delete value 2")
	linkedList.delfirstVal(2)
	linkedList.printList()

}
