package list

import "fmt"


type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

// InsertAtEnd adds a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(value int) {
    newNode := &Node{data: value, next: nil}
    if ll.head == nil {
        ll.head = newNode
        return
    }
    temp := ll.head
    for temp.next != nil {
        temp = temp.next
    }
    temp.next = newNode
}

// InsertAtBeginning adds a new node at the start of the list
func (ll *LinkedList) InsertAtBeginning(value int) {
    newNode := &Node{data: value, next: ll.head}
    ll.head = newNode
}

// DeleteNode deletes the first occurrence of a node with a given value
func (ll *LinkedList) DeleteNode(value int) {
    if ll.head == nil {
        return
    }
    if ll.head.data == value {
        ll.head = ll.head.next
        return
    }
    temp := ll.head
    for temp.next != nil && temp.next.data != value {
        temp = temp.next
    }
    if temp.next != nil {
        temp.next = temp.next.next
    }
}

func (ll *LinkedList) Search(value int) bool {
    temp := ll.head
    for temp != nil {
        if temp.data == value {
            return true
        }
        temp = temp.next
    }
    return false
}

func (ll *LinkedList) Display() {
    temp := ll.head
    for temp != nil {
        fmt.Printf("%d -> ", temp.data)
        temp = temp.next
    }
    fmt.Println("nil")
}

func SingleList() {
    ll := LinkedList{}

    ll.InsertAtEnd(10)
    ll.InsertAtEnd(20)
    ll.InsertAtEnd(30)
    ll.Display()

    ll.InsertAtBeginning(5)
    ll.Display()

    ll.DeleteNode(20)
    ll.Display()

    fmt.Println("Search 10:", ll.Search(10))
    fmt.Println("Search 50:", ll.Search(50))
}
