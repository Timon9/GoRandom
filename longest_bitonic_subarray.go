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

func solveLongestBitonicSubarray(input []int) {
	fmt.Println("LongestBitonicSubarray")

	// Build ascending auxiliary array
	a := make([]int, len(input))
	c := 0
	for i := len(input) - 2; i >= 0; i-- {
		if input[i+1] > input[i] {
			c++
		} else {
			c = 0
		}
		a[i] = c
	}

	d := make([]int, len(input))
	d[0] = 0
	c = 0
	for i := len(input) - 2; i >= 0; i-- {
		if input[i+1] < input[i] {
			c++
		} else {
			c = 0
		}
		d[i] = c
	}

	max_length := 1
	m := make(map[int]int)

	for i := 0; i < len(input)-1; i++ {
		v := a[i] + d[i+a[i]]
		if v > max_length {
			max_length = v
			m[0] = v         // value
			m[1] = input[i]  // current value input
			m[2] = i         // start highest bitonic array
			m[3] = i + v + 1 // end highest bitonic array
		}
	}

	fmt.Println("Result", input[m[2]:m[3]])

}

func LongestBitonicSubarray() {

	solveLongestBitonicSubarray([]int{3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4})
	solveLongestBitonicSubarray([]int{-5, -1, 0, 2, 1, 6, 5, 4, 2})
}
