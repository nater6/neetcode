package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	results := []struct{
		got bool
		want bool 
	}{
		{containsDuplicate([]int{1,2,3,1}),true},
		{containsDuplicate([]int{1,2,3,4}),false},
		{containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}),true},
	}
	for _, v := range results {
		if v.got != v.want {
			t.Errorf("Got: %v Want: %v", v.got, v.want)
		}
	}
}
