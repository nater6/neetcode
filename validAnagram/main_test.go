package main

import "testing"
func TestIsAnagram(t *testing.T) {
	results := []struct {
		got bool
		want bool
	} {
		{isAnagram("anagram","nagaram"), true},
		{isAnagram("car","rat"), false},
	}
	for _, r := range results {
		if r.got != r.want {
			t.Errorf("Got: %v Want: %v", r.got, r.want)
		}
	}
}