package main

import "strings"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	//Loop through the slice that we are checking
	for _, word := range strs {

		var hasBeenAdded bool
		//Go through all the anagrams found already
		for i, differentAnagrams := range result {
			//Get the first word in each anagram group
			anagramLoop: for _, anagram := range differentAnagrams {
				//Check if the words are a different length, or if one is empty and the other is not
				isSameLength := len(word) == len(anagram)
				isEmptyString :=  anagram == "" 
				isSameWord := anagram == word
				if !isSameLength || isEmptyString && !isSameWord {
					break
				}
				//check each character of the anagram and make sure it is present the same amount of times in both words
				for _, c := range anagram {
					hasSameCount := strings.Count(word, string(c)) == strings.Count(anagram, string(c))
					if !hasSameCount {
						break anagramLoop
					}
				}
				//Add to anagramcontainer[i]
				result[i] = append(result[i], word)
				hasBeenAdded = true
				break
			}
		}
		//If the word wasn't added to any anagram group create a new anagram group containing the word
		if !hasBeenAdded {
			result = append(result, []string{word})
		}
	}
    return result
}