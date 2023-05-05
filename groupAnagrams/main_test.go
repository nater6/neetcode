package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct{
		expected [][]string
		result [][]string
	}{
		{groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}), [][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}},
		{groupAnagrams([]string{""}), [][]string{{""}}},
		{groupAnagrams([]string{"a"}), [][]string{{"a"}}},
	}
	for _, test := range testCases {
		if !reflect.DeepEqual(test.expected, test.result) {
			t.Errorf("Expected != Result: %v", test)
		}
	}
}
