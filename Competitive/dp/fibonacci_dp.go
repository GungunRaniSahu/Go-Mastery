package dp

import "fmt"


func fibonacciRecursive(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func DP() {
    n := 10
    fmt.Printf("Fibonacci of %d (Recursive): %d\n", n, fibonacciRecursive(n))
}
