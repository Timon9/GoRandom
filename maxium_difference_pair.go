package main

import "fmt"

/*

Given an integer array, find the maximum difference between two elements in it such that the smaller element appears
before the larger element. If no such pair exists, return any negative number.

Input : [2, 7, 9, 5, 1, 3, 5]
Output: 7
Explanation: The pair with the maximum difference is (2, 9).

Input : [5, 4, 3, 2, 1]
Output: -1 (or any other negative number)
Explanation: No pair with the maximum difference exists.

*/

func solveMaximumDifferencePair(input []int) int {
	fmt.Println("MaximumDifferencePair")

	count := len(input)

	l := input[0]
	d := -1

	for i := 0; i < count; i++ {
		if input[i] < l {
			l = input[i]
		}

		if i != 0 && input[i]-l > d {
			d = input[i] - l
		}

	}
	if d == 0 {
		return -1
	}
	return d
}

func MaximumDifferencePair() {
	firstResult := solveMaximumDifferencePair([]int{2, 7, 9, 5, 1, 3, 5})
	secondResult := solveMaximumDifferencePair([]int{5, 4, 3, 2, 1})
	thirdResult := solveMaximumDifferencePair([]int{2, 7, 9, 5, 1, 10, 5})

	if firstResult == 7 && secondResult == -1 && thirdResult == 9 {
		fmt.Println("Succes! Completed challange.")
	} else {
		fmt.Println("Failed.")
	}

	fmt.Println("\nfirstResult: Expected 7, got", firstResult)
	fmt.Println("secondResult: Expected -1, got", secondResult)
	fmt.Println("thirdResult: Expected 9, got", thirdResult)

}
