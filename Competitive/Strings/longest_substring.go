package strings

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	maxLength, left := 0, 0

	for right, char := range s {
		if prevIndex, found := charIndex[char]; found && prevIndex >= left {
			left = prevIndex + 1
		}
		charIndex[char] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

func Longest() {
	str := "abcabcbb"
	fmt.Println("Length of longest substring:", lengthOfLongestSubstring(str))
}
