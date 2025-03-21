package main

import (
	"fmt"
)


func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid 
		} else if arr[mid] < target {
			left = mid + 1 
		} else {
			right = mid - 1 
		}
	}

	return -1 
}

func main() {
	arr := []int{3, 9, 10, 27, 38, 43, 82} 
	target := 27

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
