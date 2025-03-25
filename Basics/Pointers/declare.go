package pointers

import "fmt"

func Declare() {
    var num int = 10
    var ptr *int = &num 

    fmt.Println("Value of num:", num)
    fmt.Println("Address of num:", &num)
    fmt.Println("Value stored in ptr:", ptr)
}
