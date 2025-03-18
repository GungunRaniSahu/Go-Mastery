package main

import "fmt"

func divide(dividend int, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)
}
