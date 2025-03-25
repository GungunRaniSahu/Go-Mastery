package stack

import "fmt"


type Stack struct {
    items []int
}


func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}


func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    top := s.items[len(s.items)-1] 
    s.items = s.items[:len(s.items)-1] 
    return top, true
}


func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false 
    }
    return s.items[len(s.items)-1], true
}


func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}


func (s *Stack) Size() int {
    return len(s.items)
}


func Stack() {
    stack := Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println("Stack size:", stack.Size())

    top, _ := stack.Peek()
    fmt.Println("Top element:", top)

    popped, _ := stack.Pop()
    fmt.Println("Popped element:", popped)

    fmt.Println("Is stack empty?", stack.IsEmpty())
}
