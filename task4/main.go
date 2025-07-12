package main

import "fmt"

/*
Task: 605. Can Place Flowers
https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75
*/

func main() {
	flowerbed := []int{1,0,0,0,0,1}
	n := 2

	fmt.Println(canPlaceFlowers(flowerbed, n))
}


func canPlaceFlowers(flowerbed []int, n int) bool {
    len := len(flowerbed)
	for i := 0; i < len; i++ {
		left := i == 0 || flowerbed[i-1] == 0
		right := i == len-1 || flowerbed[i+1] == 0
		if left && right && flowerbed[i] == 0 {
			n -= 1
			flowerbed[i] = 1
		}
	}
	return n <= 0
}