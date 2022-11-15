package main

import (
	"fmt"
	"math"
)

/*
Given an integer array, find the maximum sum among all its subarrays.

Input : [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
Explanation: The maximum sum subarray is [4, -1, 2, 1]

Input : [-7, -3, -2, -4]
Output: -2
Explanation: The maximum sum subarray is [-2]

Input : [-2, 2, -1, 2, 1, 6, -10, 6, 4, -8]
Output: 10
Explanation: The maximum sum subarray is [2, -1, 2, 1, 6] or [6, 4] or [2, -1, 2, 1, 6, -10, 6, 4]
*/
func solveMaximumSumSubarray(input []int) int {
	m := math.MinInt
	c := 0

	for i := 0; i < len(input); i++ {
		v := input[i]

		c = c + v

		if c > m {
			m = c
		}

		// Two negatives can never be higher than single negative. So if its lower than 0, reset and try to find a new subarray
		if c < 0 {
			c = 0
		}

	}
	return m
}

func MaximumSumSubarray() {
	fmt.Println("MaximumSumSubarray")
	fmt.Println(solveMaximumSumSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(solveMaximumSumSubarray([]int{-7, -3, -2, -4}))
	fmt.Println(solveMaximumSumSubarray([]int{-2, 2, -1, 2, 1, 6, -10, 6, 4, -8}))

}
