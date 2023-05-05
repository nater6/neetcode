package main

import "strings"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	for _, word := range strs {
		var hasBeenAdded bool
		for i, anagramContainer := range result {
			anagramLoop: for _, anagram := range anagramContainer {
				for _, c := range anagram {
					if !strings.Contains(word, string(c)){
						break anagramLoop
					}
				}
				//Add to anagramcontainer[i]
				result[i] = append(result[i], word)
				hasBeenAdded = true
				break
			}
		}
		if !hasBeenAdded {
			result = append(result, []string{word})
		}
	}
    return result
}