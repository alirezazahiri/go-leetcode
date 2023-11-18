package main

import "fmt"

func majorityElement(nums []int) int {
	hashMap := make(map[int]int)
	maxCount, maxNum := 0, 0
	newCount := 0
	for i := 0; i < len(nums); i++ {
		_, ok := hashMap[nums[i]]

		if ok {
			hashMap[nums[i]]++
		} else {
			hashMap[nums[i]] = 1
		}

		newCount = hashMap[nums[i]]

		if maxCount < newCount {
			maxCount = newCount
			maxNum = nums[i]
		}
	}

	return maxNum
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}
