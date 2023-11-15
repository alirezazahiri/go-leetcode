package main

import (
	"fmt"
)

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	expandFromCenter := func(left, right int) string {

		for left >= 0 && right < len(s) && s[left] == s[right] {
			left -= 1
			right += 1
		}
		return s[left+1 : right]
	}

	var maxStr string = s[0:1]

	for i := 0; i < len(s); i++ {
		odd := expandFromCenter(i, i)
		even := expandFromCenter(i, i+1)

		if len(odd) > len(maxStr) {
			maxStr = odd

		}
		if len(even) > len(maxStr) {
			maxStr = even

		}
	}

	return maxStr
}

func main() {
	fmt.Println(longestPalindrome("baabad"))
	fmt.Println(longestPalindrome("dabab"))
}
