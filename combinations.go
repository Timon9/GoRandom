package main

import (
	"fmt"
)

/*
Given an integer array, find all distinct combinations of a given length `k`.
The solution should return a set containing all the distinct combinations, while preserving the relative order of elements as they appear in the array.

Input : [2, 3, 4], k = 2
Output: {[2, 3], [2, 4], [3, 4]}

Input : [1, 2, 1], k = 2
Output: {[1, 2], [1, 1], [2, 1]}

Input : [1, 1, 1], k = 2
Output: {[1, 1]}

Input : [1, 2, 3], k = 4
Output: {}

Input : [1, 2], k = 0
Output: {[]}
*/

// combination returns all distinct combinations of length k in an integer array
func combination(arr []int, k int) [][]int {
	// result to store the final combinations
	var result [][]int

	// recursive function to generate combinations
	var combinationUtil func(arr []int, k int, start int, subset []int)
	combinationUtil = func(arr []int, k int, start int, subset []int) {
		// if length of current subset is equal to k, add it to the result
		if len(subset) == k {
			result = append(result, append([]int(nil), subset...))
			return
		}

		// generate combinations by adding current element to the subset and calling the function recursively
		for i := start; i < len(arr); i++ {
			// skip duplicate elements
			if i > start && arr[i] == arr[i-1] {
				continue
			}
			combinationUtil(arr, k, i+1, append(subset, arr[i]))
		}
	}

	// call the function to generate combinations
	combinationUtil(arr, k, 0, []int{})

	return result
}

func Combinations() {
	fmt.Println("Combinations")

	arr := []int{1, 2, 1}
	k := 2
	fmt.Println(combination(arr, k))

	// solveCombinations([]int{2, 3, 4}, 2)
	// solveCombinations([]int{1, 2, 1}, 2)
	// solveCombinations([]int{1, 1, 1}, 2)
	// solveCombinations([]int{1, 2, 3}, 4)
	// solveCombinations([]int{1, 2, 3, 5, 2, 1}, 3)
	// //solveCombinations([]int{1, 2}, 0)
}
