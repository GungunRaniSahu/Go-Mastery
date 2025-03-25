package main

import "fmt"

func updateValue(n *int) {
    *n = 100
}

func main() {
    num := 25
    fmt.Println("Before:", num)
    updateValue(&num)
    fmt.Println("After:", num)
}
