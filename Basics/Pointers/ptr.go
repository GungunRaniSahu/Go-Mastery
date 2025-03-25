package pointers

import "fmt"

func Ptr() {
    var num int = 20
    var ptr *int = &num

    fmt.Println("Value stored in ptr:", *ptr)
}
