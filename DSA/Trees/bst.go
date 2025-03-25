package tree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(data int) {
	bst.root = insertRec(bst.root, data)
}

func insertRec(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	}
	if data < root.data {
		root.left = insertRec(root.left, data)
	} else {
		root.right = insertRec(root.right, data)
	}
	return root
}

func (bst *BST) InOrder() {
	inOrderRec(bst.root)
	fmt.Println()
}

func inOrderRec(root *Node) {
	if root != nil {
		inOrderRec(root.left)
		fmt.Print(root.data, " ")
		inOrderRec(root.right)
	}
}

func (bst *BST) Search(data int) bool {
	return searchRec(bst.root, data)
}

func searchRec(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	if data < root.data {
		return searchRec(root.left, data)
	}
	return searchRec(root.right, data)
}

func BinarySearchTree() {
	bst := BST{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	bst.InOrder()

	fmt.Println(bst.Search(40))
	fmt.Println(bst.Search(100))
}
