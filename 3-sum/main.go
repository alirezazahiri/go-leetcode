package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triple := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triple)

					found := false
					for _, t := range res {
						if t[0] == triple[0] && t[1] == triple[1] && t[2] == triple[2] {
							found = true
							break
						}
					}
					if !found {
						res = append(res, triple)
					}
				}
			}
		}
	}

	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // Skip same numbers
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}

func next(nums []int, l int, r int) (int, int) {
	l++
	r--
	for l < r && nums[l] == nums[l-1] {
		l++
	}
	for l < r && nums[r] == nums[r+1] {
		r--
	}
	return l, r
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
