package main

import "fmt"

func main() {
    var num int = 20
    var ptr *int = &num

    fmt.Println("Value stored in ptr:", *ptr)
}
