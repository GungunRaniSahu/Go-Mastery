package list

import (
	"fmt"
)

type DNode struct {
	data int
	prev *DNode
	next *DNode
}

type DoublyLinkedList struct {
	head *DNode // Use DNode instead of Node
}

// InsertAtEnd adds a node at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &DNode{data: data}

	if dll.head == nil {
		dll.head = newNode
		return
	}

	temp := dll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	newNode.prev = temp
}

// InsertAtBeginning adds a node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &DNode{data: data, next: dll.head}
	if dll.head != nil {
		dll.head.prev = newNode
	}
	dll.head = newNode
}

// DeleteNode removes a node by value
func (dll *DoublyLinkedList) DeleteNode(data int) {
	if dll.head == nil {
		return
	}

	temp := dll.head

	if temp.data == data {
		dll.head = temp.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		return
	}

	for temp != nil && temp.data != data {
		temp = temp.next
	}

	if temp == nil {
		return
	}

	if temp.next != nil {
		temp.next.prev = temp.prev
	}
	if temp.prev != nil {
		temp.prev.next = temp.next
	}
}

// DisplayForward prints the list from head to tail
func (dll *DoublyLinkedList) DisplayForward() {
	temp := dll.head
	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

// DisplayBackward prints the list from tail to head
func (dll *DoublyLinkedList) DisplayBackward() {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}

	temp := dll.head
	for temp.next != nil {
		temp = temp.next
	}

	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		temp = temp.prev
	}
	fmt.Println("nil")
}

func DoublyList() {
	dll := &DoublyLinkedList{}

	dll.InsertAtEnd(10)
	dll.InsertAtEnd(20)
	dll.InsertAtEnd(30)
	dll.InsertAtBeginning(5)

	fmt.Println("Forward Traversal:")
	dll.DisplayForward()

	fmt.Println("Backward Traversal:")
	dll.DisplayBackward()

	fmt.Println("Deleting node with value 20")
	dll.DeleteNode(20)
	dll.DisplayForward()
}
