package main

func twoSum(nums []int, target int) []int {
	//Iterate through each index of nums
	for i, n := range nums {
		//Iterate again skip the current index
		for j:=i+1; j< len(nums); j++ {
			//check if result = target
			if j != i && n+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return []int{}
}
