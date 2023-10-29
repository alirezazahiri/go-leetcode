package main

import "fmt"

func plusOne(digits []int) []int {
	last := len(digits) - 1

	if last == -1 {
		return []int{1}
	}

	digits[last]++
	for i := last; i > 0 && digits[i] == 10; i-- {
		digits[i] = 0
		digits[i-1]++
	}

	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(plusOne(arr))
	arr = []int{1, 2, 9}
	fmt.Println(plusOne(arr))
	arr = []int{1, 2, 9, 9}
	fmt.Println(plusOne(arr))
	arr = []int{9}
	fmt.Println(plusOne(arr))
	arr = []int{}
	fmt.Println(plusOne(arr))
	arr = []int{9, 9}
	fmt.Println(plusOne(arr))
}
