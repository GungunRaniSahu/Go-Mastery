package functions

import "fmt"

func multiply(a int, b int) int {
    return a * b
}

func Return() {
    result := multiply(4, 5)
    fmt.Println("Product:", result)
}
