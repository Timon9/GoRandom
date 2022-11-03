package main

import "fmt"

/*

The Longest Bitonic Subarray (LBS) problem is to find a contiguous subarray of a given sequence in which the subarray’s elements are first sorted in increasing order, then in decreasing order, and the subarray is as long as possible.

Input : [3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4]
Output: [4, 5, 9, 10, 8, 5, 3]

In case the multiple bitonic subarrays of maximum length exists, the solution can return any one of them.

Input : [-5, -1, 0, 2, 1, 6, 5, 4, 2]
Output: [-5, -1, 0, 2, 1] or [1, 6, 5, 4, 2]

*/

func LongestBitonicSubarray() {
	fmt.Println("LongestBitonicSubarray")
	input := []int{3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4}

	// Build ascending auxiliary array
	a := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		a[i] = 1
	}
	fmt.Println(a)
}
