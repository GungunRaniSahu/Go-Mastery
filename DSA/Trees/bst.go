package tree

import "fmt"

type BSTNode struct {  
    data  int  
    left  *BSTNode  
    right *BSTNode  
}


type BST struct {
	root *BSTNode
}

func (bst *BST) Insert(data int) {
	bst.root = InsertRec(bst.root, data)
}

func InsertRec(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return &BSTNode{data: data}
	}
	if data < root.data {
		root.left = InsertRec(root.left, data)
	} else {
		root.right = InsertRec(root.right, data)
	}
	return root
}

func (bst *BST) InOrder() {
	InOrderRec(bst.root)
	fmt.Println()
}

func InOrderRec(root *BSTNode) {
	if root != nil {
		InOrderRec(root.left)
		fmt.Print(root.data, " ")
		InOrderRec(root.right)
	}
}

func (bst *BST) Search(data int) bool {
	return SearchRec(bst.root, data)
}

func SearchRec(root *BSTNode, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	if data < root.data {
		return SearchRec(root.left, data)
	}
	return SearchRec(root.right, data)
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
