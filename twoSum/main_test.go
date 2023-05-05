package main

import (
	"reflect"
	"testing"
)


func TestTwoSum(t *testing.T) {
	testCases := []struct{
		result []int
		expected []int
	}{
		{twoSum([]int{2,7,11,15}, 9),[]int{0,1}},
		{twoSum([]int{3,2,4},6),[]int{1,2}},
		{twoSum([]int{3,3}, 6),[]int{0,1}},
	}

	for _, test := range testCases {
		if !reflect.DeepEqual(test.expected, test.result) {
			t.Errorf("Expected != Result: %v", test)
		}
	} 
	
}
