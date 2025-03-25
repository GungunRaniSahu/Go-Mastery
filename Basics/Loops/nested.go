package loops

import "fmt"

func Nested() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("(%d, %d) ", i, j)
        }
        fmt.Println()
    }
}
