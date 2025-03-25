package pointers

import "fmt"

func Nil() {
    var ptr *int 
    fmt.Println("Pointer value:", ptr)

    if ptr == nil {
        fmt.Println("Pointer is nil")
    }
}
