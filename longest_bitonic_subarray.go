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
	fmt.Println(input)

	// Build ascending auxiliary array
	a := make([]int, len(input))
	a[0] = 1
	c := 1
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			c++
		} else {
			c = 1
		}
		a[i] = c
	}
	fmt.Println("ASC", a)

	// Build descending auxiliary array
	d := make([]int, len(input))
	d[len(input)-1] = 1
	c = 1
	for i := len(input) - 2; i >= 0; i-- {
		if input[i+1] > input[i] {
			c++
		} else {
			c = 1
		}
		d[i] = c
	}

	max_length := 1
	m := make(map[int]int)

	for i := 0; i < len(input); i++ {
		v := a[i] + d[i]
		if v > max_length {
			max_length = v
			m[0] = v            // value
			m[1] = input[i]     // current value input
			m[2] = i - d[i] + 1 // start highest bitonic array
			m[3] = i + a[i] - 1 // end highest bitonic array
		}
	}
	//	fmt.Println(m)

	fmt.Println("DESC", d)
	fmt.Println("Result", input[m[2]:m[3]])
}
