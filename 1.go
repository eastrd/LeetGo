// Brute force by looping n^2 and compare if the two indices add up to the target

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for idx1 := range nums {
		for idx2 := range nums {
			if idx1 != idx2 && nums[idx1]+nums[idx2] == target {
				return []int{idx1, idx2}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
