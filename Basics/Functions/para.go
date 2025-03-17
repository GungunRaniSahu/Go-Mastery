package main

import "fmt"

func add(a int, b int) {
    sum := a + b
    fmt.Println("Sum:", sum)
}

func main() {
    add(5, 3)
}
