package main

type Node struct {
	data      int
	leftNode  *Node // < data
	rightNode *Node // >= data
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) addNode(data int) {
	newNode := &Node{data: data}
	if bt.root == nil {
		bt.root = newNode
		return
	}
	curr := bt.root
	for {
		if newNode.data < curr.data {
			if curr.leftNode == nil {
				curr.leftNode = newNode
				return
			}
			curr = curr.leftNode
		} else {
			if curr.rightNode == nil {
				curr.rightNode = newNode
				return
			}
			curr = curr.rightNode
		}
	}
}

func (bt *BinaryTree) printTreeFromNode(n *Node) {
	if n == nil {
		return
	}
	println(n.data)
	bt.printTreeFromNode(n.leftNode)
	bt.printTreeFromNode(n.rightNode)
}

func (bt *BinaryTree) printFullTree() {
	bt.printTreeFromNode(bt.root)
}

func main() {

	binaryTree := &BinaryTree{}

	binaryTree.addNode(5)
	binaryTree.addNode(4)
	binaryTree.addNode(6)
	binaryTree.addNode(7)
	binaryTree.addNode(1)

	binaryTree.printFullTree()
}
