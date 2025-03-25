package pointers

import "fmt"

func modify(arr *[3]int) {
    (*arr)[0] = 100
}

func Array() {
    numbers := [3]int{1, 2, 3}
    fmt.Println("Before:", numbers)

    modify(&numbers)
    fmt.Println("After:", numbers)
}
