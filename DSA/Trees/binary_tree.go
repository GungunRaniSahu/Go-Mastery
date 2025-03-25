package tree

import (
	"fmt"
)


type Node struct {
	data  int
	left  *Node
	right *Node
}


type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) Insert(data int) {
	bt.root = insertRec(bt.root, data)
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


func (bt *BinaryTree) InOrderTraversal() {
	inOrderRec(bt.root)
	fmt.Println()
}


func inOrderRec(root *Node) {
	if root != nil {
		inOrderRec(root.left)
		fmt.Print(root.data, " ")
		inOrderRec(root.right)
	}
}


func (bt *BinaryTree) PreOrderTraversal() {
	preOrderRec(bt.root)
	fmt.Println()
}


func preOrderRec(root *Node) {
	if root != nil {
		fmt.Print(root.data, " ")
		preOrderRec(root.left)
		preOrderRec(root.right)
	}
}


func (bt *BinaryTree) PostOrderTraversal() {
	postOrderRec(bt.root)
	fmt.Println()
}


func postOrderRec(root *Node) {
	if root != nil {
		postOrderRec(root.left)
		postOrderRec(root.right)
		fmt.Print(root.data, " ")
	}
}


func (bt *BinaryTree) Search(data int) bool {
	return searchRec(bt.root, data)
}


func searchRec(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	} else if data < root.data {
		return searchRec(root.left, data)
	} else {
		return searchRec(root.right, data)
	}
}

func BinaryTrees() {
	bt := &BinaryTree{}
	bt.Insert(50)
	bt.Insert(30)
	bt.Insert(70)
	bt.Insert(20)
	bt.Insert(40)
	bt.Insert(60)
	bt.Insert(80)

	fmt.Println("In-order Traversal:")
	bt.InOrderTraversal()

	fmt.Println("Pre-order Traversal:")
	bt.PreOrderTraversal()

	fmt.Println("Post-order Traversal:")
	bt.PostOrderTraversal()


	searchValue := 40
	if bt.Search(searchValue) {
		fmt.Printf("Value %d found in the tree\n", searchValue)
	} else {
		fmt.Printf("Value %d not found in the tree\n", searchValue)
	}
}
