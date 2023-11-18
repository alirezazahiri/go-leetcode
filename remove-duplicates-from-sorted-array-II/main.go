package main

import "fmt"

func removeDuplicates(nums []int) int {
    counter := 0
    for _, num := range(nums) {
        if (counter <= 1 || nums[counter-2] != num) {
            nums[counter] = num
            counter++
        }
    }

    return counter
}

func main() {
	arr:=[]int{1,1,1,2,2,3}
	removeDuplicates(arr)
	fmt.Println(arr)
}