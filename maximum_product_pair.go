package main

import "fmt"

/*

Given an integer array, find a pair with the maximum product in it.

Each input can have multiple solutions. The output should match with either one of them.

Input : [-10, -3, 5, 6, -2]
Output: (-10, -3) or (-3, -10) or (5, 6) or (6, 5)

Input : [-4, 3, 2, 7, -5]
Output: (3, 7) or (7, 3)

If no pair exists, the solution should return null.

Input : [1]
Output: null

*/

func sort(input []int) []int {

	var r = []int{}

	piv := 1
	c := 0
	for i := 0; i < len(input); i++ {
		v := input[i]
		if v > piv {
			//Append after pivot
			r = append(r, v)
		} else {
			//Prepppend before pivot
			r = append([]int{v}, r...)
			c++
		}
	}
	fmt.Printf("Current pivot point:%d\n", len(r)-c)
	return r
}

func maximumProductPair(input []int) []int {
	fmt.Println("MaximumProductPair")

	maxPair := []int{0, 0}

	/*
		Idea: Sort the integers, select the top end 2 and compare with the bottom end 2 (-x*-x=pos)
	*/
	fmt.Println("====")
	fmt.Println("Before sort:")
	fmt.Println(input)

	input = sort(input)
	fmt.Println("After sort:")
	fmt.Println(input)

	fmt.Println("====")

	return maxPair

}

func StartMaximumProductPair() {
	fmt.Println(maximumProductPair([]int{-10, -3, 5, 6, -2}))
	//	fmt.Println(maximumProductPair([]int{-4, 3, 2, 7, -5}))
	//	fmt.Println(maximumProductPair([]int{1}))

}
