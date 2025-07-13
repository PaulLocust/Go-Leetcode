package main

import (
	"fmt"
	"strings"
)

/*
Task: 345. Reverse Vowels of a String
https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75
Так как s будет содержать только ASCII символы то использую byte, иначе rune
*/

func main() {
	s := "Yo! Bottoms up, U.S. Motto, boy!" //
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	bytes_str := []byte(s) // Преобразуем в срез рун

	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer <= rightPointer {
		if contains(vowels, s[leftPointer]) && contains(vowels, s[rightPointer]) {
			bytes_str[leftPointer], bytes_str[rightPointer] = bytes_str[rightPointer], bytes_str[leftPointer]
			leftPointer++
			rightPointer--

		} else if contains(vowels, s[leftPointer]) && !contains(vowels, s[rightPointer]) {
			rightPointer--
		} else {
			leftPointer++
		}
	}
	return string(bytes_str)
}

func contains(alphabet string, letter uint8) bool {
	return strings.Contains(alphabet, string(letter))
}
