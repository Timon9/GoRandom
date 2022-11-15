package main

import "fmt"

/*
Given a circular integer array, find a contiguous subarray with the largest sum in it.

https://www.techiedelight.com/maximum-sum-circular-subarray/

Input : [2, 1, -5, 4, -3, 1, -3, 4, -1]
Output: 6
Explanation: Subarray with the largest sum is [4, -1, 2, 1] with sum 6.

Input : [8, -7, -3, 5, 6, -2, 3, -4, 2]
Output: 18
Explanation: Subarray with the largest sum is [5, 6, -2, 3, -4, 2, 8] with sum 18.
*/
func solveMaximumSumCircularSubarray(input []int) {
	input = append(input, input...) // Circler
	fmt.Println(input)
	count := len(input)

	m := 0
	c := 0
	s := 0

	for i := 0; i < count; i++ {
		v := input[i]

		c = c + v

		if c > m && s < count/2 {
			m = c
		}

		s++

		if c < 0 {
			c = 0
			s = 0
		}

	}

	fmt.Println(m)
}
func MaximumSumCircularSubarray() {
	fmt.Println("MaximumSumCircularSubarray")
	solveMaximumSumCircularSubarray([]int{2, 1, -5, 4, -3, 1, -3, 4, -1})
	solveMaximumSumCircularSubarray([]int{8, -7, -3, 5, 6, -2, 3, -4, 2})
}
