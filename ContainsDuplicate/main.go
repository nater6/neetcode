package main

import "sort"

func containsDuplicate(nums []int) bool {
	// Nested Slice loop
	sort.Ints(nums)
	for i, x := range nums {
		if i > 0 && nums[i-1] == x {
			return true
		}
	}
	return false
}
