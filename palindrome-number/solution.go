package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {
         return false;
    }
	strInt := strconv.Itoa(x)
	for i := 0; i < len(strInt)/2; i++ {
		if strInt[i] != strInt[len(strInt)-1-i] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(isPalindrome(151))
	fmt.Println(isPalindrome(-151))
}