package main

import "fmt"

/*

Given a binary array, find the index of 0 to be replaced with 1 to get the maximum length sequence of continuous ones. The solution should return the index of first occurence of 0, when multiple continuous sequence of maximum length is possible.

Input : [0, 0, 1, 0, 1, 1, 1, 0, 1, 1]
Output: 7
Explanation: Replace index 7 to get the continuous sequence of length 6 containing all 1’s.

Input : [0, 1, 1, 0, 0]
Output: 0
Explanation: Replace index 0 or 3 to get the continuous sequence of length 3 containing all 1’s. The solution should return the first occurence.

Input : [1, 1]
Output: -1
Explanation: Invalid Input (all 1’s)

*/

func solveMaximumContinuousSequence(input []int) int {

	hm := []int{0, 0, 0} // Max count, max position, current count

	for i := 0; i < len(input); i++ {
		v := input[i]

		if v == 1 {
			hm[2]++
		} else {
			hm[2] = 0
		}

		if hm[2] > hm[0] {
			hm[2] = hm[0]
			hm[1] = i - 2
		}
	}

	return hm[1]

}

func MaximumContinuousSequence() {
	fmt.Println("MaximumContinuousSequence")
	r := solveMaximumContinuousSequence([]int{0, 0, 1, 0, 1, 1, 1, 0, 1, 1})
	if r == 7 {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failed, expected 7 got ", r)
	}

	r2 := solveMaximumContinuousSequence([]int{0, 1, 1, 0, 0})
	if r2 == 0 {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failed, expected 0 got ", r2)
	}

	r3 := solveMaximumContinuousSequence([]int{1, 1})
	if r3 == -1 {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failed, expected -1 got ", r3)
	}
}
