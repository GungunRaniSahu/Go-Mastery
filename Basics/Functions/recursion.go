package functions

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func Recursion() {
    fmt.Println("Factorial of 5:", factorial(5))
}
