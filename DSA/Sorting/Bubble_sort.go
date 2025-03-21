package main

import "fmt"


func QuickSort(arr []int) 
{
	if len(arr) <= 1 {
		return
	}
	

	pivotIndex := partition(arr)
	

	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}


func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	QuickSort(arr)

	fmt.Println("Sorted array:", arr)
}
