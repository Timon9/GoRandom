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
	for i := 0; i < len(input); i++ {
		v := input[i]

		fmt.Println(v)

	}
}

func ReplaceArrayElements() {
	solveReplaceArrayElements([]int{1, 2, 3, 4, 5})
	solveReplaceArrayElements([]int{5, 3, 4, 2, 6, 8})

}
