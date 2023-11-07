package main 

import "fmt"

func pow(x, n int) int {
	res := 1
    for i:=0;i<n;i++ {
        res *= x
    }
    return res
}

func numIntoDigitsArray(num int) []int {
	digits := []int{}

    for ;num != 0; {
		v := num%10
        digits = append(digits, v)
		num /= 10
    }

	return digits
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

    digits := numIntoDigitsArray(x)
    reversedNum := 0

    for i,digit := range(digits) {
        reversedNum += pow(10, len(digits) - i - 1) * digit
    }

    return reversedNum == x
}

func main() {
	fmt.Println(isPalindrome(151))
	fmt.Println(isPalindrome(-151))
}