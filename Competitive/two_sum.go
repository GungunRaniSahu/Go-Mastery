package array

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := hash[complement]; ok {
			return []int{idx, i}
		}
		hash[num] = i
	}
	return []int{}
}

func Two_sum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Indices:", result) 
}
