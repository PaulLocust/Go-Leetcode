package main

/*
Task: 1071. Greatest Common Divisor of Strings
https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75
*/

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	fmt.Println(gcdOfStrings(str1, str2))
	fmt.Println(gcdOfStringsEuclid(str1, str2))
}

// Первый вариант реализации
func gcdOfStrings(str1 string, str2 string) string {
	buf := ""
	if len(str1) < len(str2) {
		buf = str1
		str1 = str2
		str2 = buf
	}
	ans := str2
	for i:=len(str2)-1; i >= 0; i-- {
		if len(str1) % len(ans) != 0 || len(str2) % len(ans) != 0{
			ans = str2[0:i]
			continue
		}
		
		if strings.Repeat(ans, len(str1) / len(ans)) == str1 && strings.Repeat(ans, len(str2) / len(ans)) == str2{
			return ans
		} else {
			ans = str2[0:i]
		}
	}
	return ""
}

// Второй вариант реализации
func gcdOfStringsEuclid(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	a, b := len(str1), len(str2)

	for b > 0 {
		a, b = b, a%b	 
	}
	return str1[:a]
}