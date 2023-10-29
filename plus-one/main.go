package main

import "fmt"

func plusOne(digits []int) []int {

	var last int = len(digits) - 1

	if last == -1 {
		return []int{1}
	}

	last_digit := digits[last] + 1
	i := 0

	for last_digit == 10 && last - i >= 0 {
		digits[last-i] = 0
		i++
		if last - i >= 0 {
			last_digit = digits[last-i] + 1
		}
	}
	if last-i >= 0 {
		digits[last-i] = last_digit
	} else {
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
