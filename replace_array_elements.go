package main

import "fmt"

/*

Given an integer array, replace each element with the product of every other element without using the division operator.

Input : [1, 2, 3, 4, 5]
Output: [120, 60, 40, 30, 24]

Input : [5, 3, 4, 2, 6, 8]
Output: [1152, 1920, 1440, 2880, 960, 720]

*/

func solveReplaceArrayElements(input []int) {

	fmt.Println("ReplaceArrayElements")

	// `left[i]` stores the product of all elements in subarray[0…i-1]
	left := make(map[int]int)
	left[0] = 1
	for i := 1; i < len(input); i++ {
		left[i] = input[i-1] * left[i-1]
	}

	// `right[i]` stores the product of all elements in subarray[i+1…n-1]
	right := make(map[int]int)
	right[len(input)-1] = 1
	for ii := len(input) - 2; ii >= 0; ii-- {
		right[ii] = input[ii+1] * right[ii+1]
	}

	// replace each element with the product of its left and right subarray
	for i := 0; i < len(input); i++ {
		input[i] = left[i] * right[i]
	}

	fmt.Println(input)
}

func ReplaceArrayElements() {
	solveReplaceArrayElements([]int{1, 2, 3, 4, 5})
	solveReplaceArrayElements([]int{5, 3, 4, 2, 6, 8})

}
