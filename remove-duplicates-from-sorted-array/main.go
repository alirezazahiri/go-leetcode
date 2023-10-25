package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i, num := range nums {
		for i+1 < len(nums) && num == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

// func removeDuplicates(nums []int) int {
// 	prev := nums[0]
// 	l := 1
// 	for _, num := range nums {
// 		if num != prev {
// 			nums[l] = num
// 			l++
// 		}
// 		prev = num
// 	}
// 	return l
// }

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
