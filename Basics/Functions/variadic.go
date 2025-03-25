package functions

import "fmt"

func sum(numbers ...int) {
    total := 0
    for _, num := range numbers {
        total += num
    }
    fmt.Println("Sum:", total)
}

func Variadic() {
    sum(1, 2, 3, 4, 5) 
}
