package queue

import "fmt"


type Queue struct {
    items []int
}


func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}


func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false 
    }
    front := q.items[0] 
    q.items = q.items[1:]
    return front, true
}


func (q *Queue) Peek() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    return q.items[0], true
}


func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}


func (q *Queue) Size() int {
    return len(q.items)
}


func Queues() {
    queue := Queue{}

    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    fmt.Println("Queue size:", queue.Size())

    front, _ := queue.Peek()
    fmt.Println("Front element:", front)

    dequeued, _ := queue.Dequeue()
    fmt.Println("Dequeued element:", dequeued)

    fmt.Println("Is queue empty?", queue.IsEmpty())
}
