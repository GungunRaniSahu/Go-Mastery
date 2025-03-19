package main

import "fmt"

func main() {
    num := 30
    ptr := &num

    fmt.Println("Before:", num)
    *ptr = 50
    fmt.Println("After:", num)
}
