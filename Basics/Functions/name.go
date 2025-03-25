package functions

import "fmt"

func rectangleDimensions(length int, breadth int) (area int, perimeter int) {
    area = length * breadth
    perimeter = 2 * (length + breadth)
    return 
}

func Name() {
    a, p := rectangleDimensions(5, 3)
    fmt.Println("Area:", a, "Perimeter:", p)
}
