package main

import "fmt"

/*
Task: 1431. Kids With the Greatest Number of Candies
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75
*/

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var ans []bool
	max_candies := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > max_candies {
			max_candies = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i] + extraCandies >= max_candies {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}
