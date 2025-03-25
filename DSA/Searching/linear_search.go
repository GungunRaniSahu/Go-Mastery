package searching

import "fmt"

func LinearSearch(arr []int, target int) int {
    for i, num := range arr  {
        if num == target  {
            return i 
        }
    }
    return -1 
}

func Linear() {
    arr := []int{10, 20, 30, 40, 50}
    target := 30

    index := LinearSearch(arr, target)

    if index != -1 {
        fmt.Printf("Element %d found at index %d\n", target, index)
    }else  {
        fmt.Printf("Element %d not found\n", target)
    }
}
