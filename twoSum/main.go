package main

func twoSum(nums []int, target int) []int {
	//Iterate through each index of nums
	for i, n := range nums {
		//Iterate again skip the current index
		for j, x := range nums {
			//check if result = target
			if j != i && n+x == target {
				return []int{i, j}
			}
		}

	}
	return []int{}
}
