package main

import (
	"fmt"
	"strings"
)

/*
Task: 151. Reverse Words in a String
https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75
*/

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	ans := ""
	stringArr := strings.Fields(s)
	for i := len(stringArr) - 1; i >= 0; i-- {
		ans += stringArr[i] + " "
	}
	return strings.TrimSpace(ans)
}
