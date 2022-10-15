package main

import "fmt"

/*

Given an array containing only 0’s, 1’s, and 2’s, sort it in linear time and using constant space.

Input : [0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0]
Output: [0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2]

*/

func solve(input []int) []int {

	var s []int
	var m []int
	var e []int

	for i := 0; i < len(input); i++ {
		v := input[i]
		switch v {
		case 0:
			s = append(s, v)
		case 1:
			m = append(m, v)
		case 2:
			e = append(e, v)

		}
	}
	//Combine the arrays

	return append(s, append(m, e...)...)
}

func DutchNationalFlagProblem() {
	fmt.Println("Dutch National Flag Problem")
	fmt.Println(solve([]int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}))
}
