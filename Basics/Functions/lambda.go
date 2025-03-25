package functions

import "fmt"

func Lambda() {
    add := func(a int, b int) int {
        return a + b
    }

    fmt.Println("Addition:", add(3, 7))
}
