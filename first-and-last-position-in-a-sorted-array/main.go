package main

import "fmt"

func searchRange(nums []int, target int) []int {
    indexes := []int{-1, -1}
	count := 0

	for i, num := range(nums) {
		if num == target  {
			if count == 0 {
				indexes[0] = i
				indexes[1] = i
				count++
			} else {
				indexes[1] = i
			}
		}	
	}

	return indexes
}

func main() {
	nums, target := []int{5,7,7,8,8,10}, 8
	fmt.Println(searchRange(nums, target))
	nums, target = []int{5,7,7,8,8,10}, 6
	fmt.Println(searchRange(nums, target))
	nums, target = []int{1}, 1
	fmt.Println(searchRange(nums, target))
}