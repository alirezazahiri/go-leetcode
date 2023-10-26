package main

import "fmt"

func filter[T any](arr *[]T, callback func(item T, index int) bool) []T {
	filteredArray := []T{}

	for index, item:= range(*arr) {
		if callback(item, index) {
			filteredArray = append(filteredArray, item)
		}
	}

	*arr = filteredArray

	return filteredArray
}

func removeElement(nums *[]int, val int) ([]int, int) {
    filteredArray := filter[int](nums, func (item, index int) bool {
		return item != val
	})
    return filteredArray, len(filteredArray)
}

func main() {
	arr := []int{0,1,2,2,3,0,4,2}
	arr, out := removeElement(&arr, 2)
	fmt.Println(arr, out)
	
	arr = []int{3, 2, 2, 3}
	new_arr, out := removeElement(&arr, 3)
	fmt.Println(new_arr, arr, out)
}