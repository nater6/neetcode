package main

func containsDuplicate(nums []int) bool {
	// Nested Array loop
	for i, x := range nums {
		for j, y := range nums {
			if i != j && x == y {
				return true
			}
		}
	}
	return false
}
