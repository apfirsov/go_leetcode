//https://leetcode.com/problems/binary-search/

package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] > target || nums[r] < target {
			return -1
		}
		c := (l + r) / 2
		if nums[c] == target {
			return c
		} else if nums[c] > target {
			r = c - 1
		} else {
			l = c + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{2, 5}, 5))
}
