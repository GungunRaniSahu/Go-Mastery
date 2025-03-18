package main

import "fmt"

func main() {
    add := func(a int, b int) int {
        return a + b
    }

    fmt.Println("Addition:", add(3, 7))
}
