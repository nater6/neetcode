package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	result_s := make(map[rune]int)
	result_t := make(map[rune]int)
	for _, x := range s {
		if _, ok := result_s[x]; ok {
			result_s[x]++
		} else {
			result_s[x] = 1
		}
	}
	for k, v := range result_s {
		fmt.Println("Key: ", k, " Value: ", v)
	}
	for _, y := range t {
		if _, ok := result_t[y]; !ok { // rune wasn't present in first string
		result_t[y] = 1
		
		} else {
			result_t[y]++
		}
	}
	return reflect.DeepEqual(result_s, result_t)
}
