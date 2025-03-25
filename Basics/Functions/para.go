package functions

import "fmt"

func add(a int, b int) {
    sum := a + b
    fmt.Println("Sum:", sum)
}

func Para() {
    add(5, 3)
}
