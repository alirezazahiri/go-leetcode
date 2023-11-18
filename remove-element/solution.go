package main

import "fmt"

func removeElement(nums []int, val int) int {
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}

	return counter
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	out := removeElement(arr, 2)
	fmt.Println("------------------------------------")
	fmt.Println(out)
	fmt.Println("------------------------------------")

	arr = []int{3, 2, 2, 3}
	out = removeElement(arr, 3)
	fmt.Println("------------------------------------")
	fmt.Println(out)
	fmt.Println("------------------------------------")
	
	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	out = removeElement(arr, 2)
	fmt.Println("------------------------------------")
	fmt.Println(arr)
	fmt.Println(out)
	fmt.Println("------------------------------------")
}
