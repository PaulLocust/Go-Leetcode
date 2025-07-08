package main

/*
Task: 1768. Merge Strings Alternately
https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
	"strings"
)

func main() {
	word1 := "ab"
	word2 := "pqrs"
	fmt.Println(mergeAlternately(word1, word2))

}

func mergeAlternately(word1 string, word2 string) string {
	
	if len(word1) < len(word2) {
		word1 = word1 + strings.Repeat(" ", len(word2) - len(word1))
	} else if len(word1) > len(word2) {
		word2 = word2 + strings.Repeat(" ", len(word1) - len(word2))
	}

	n := len(word1) + len(word2)
	storage := make([]rune, n)
	pointer1 := 0
	pointer2 := 1

	for _, c := range word1 {
		storage[pointer1] = c
		pointer1 += 2
	}
	for _, c := range word2 {
		storage[pointer2] = c
		pointer2 += 2
	}
	
	res := strings.ReplaceAll(string(storage), " ", "")
	return res
}
