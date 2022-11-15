package main

import (
	"fmt"
	"math"
)

/*
Given an integer array, find a contiguous subarray within it that has the maximum sum.

Input : [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: [4, -1, 2, 1]

Input : [-7, -3, -2, -4]
Output: [-2]

In case multiple subarrays exists with the largest sum, the solution can return any one of them.

Input : [-2, 2, -1, 2, 1, 6, -10, 6, 4, -8]
Output: [2, -1, 2, 1, 6] or [6, 4] or [2, -1, 2, 1, 6, -10, 6, 4]
*/
func solveMaximumSumSubarrayII(input []int) (int, []int) {
	count := len(input)
	m := math.MinInt
	c := 0
	startPos := 0
	endPos := 0
	b := 0

	for i := 0; i < count; i++ {
		v := input[i]

		if c == 0 {
			b = i
		}

		c = c + v
		if c > m {
			m = c
			endPos = i
			startPos = b
		}
		// Negative sum, reset the count, its not getting higher than this because one negative will be higher than two negatives.
		if c < 0 {
			c = 0
		}

	}

	endPos++
	fmt.Println("Positions:", startPos, endPos)
	slice := input[startPos:endPos]
	return m, slice
}

func MaximumSumSubarrayII() {
	fmt.Println("MaximumSumSubarrayII")

	fmt.Println(solveMaximumSumSubarrayII([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(solveMaximumSumSubarrayII([]int{-7, -3, -2, -4}))
	fmt.Println(solveMaximumSumSubarrayII([]int{-2, 2, -1, 2, 1, 6, -10, 6, 4, -8}))
}
