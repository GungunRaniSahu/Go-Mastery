package loops

import "fmt"

func ForInRange() {
    nums := []int{10, 20, 30}
    for index, value := range nums {
        fmt.Println("Index:", index, "Value:", value)
    }
}
