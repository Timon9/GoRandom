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

https://www.geeksforgeeks.org/quick-sort/
https://nl.wikipedia.org/wiki/Quicksort
https://www.techiedelight.com/find-maximum-product-two-integers-array/
https://www.techiedelight.com/?problem=MaximumProductPair

*/

func partition(input []int) ([]int, int, int) {
	piv := input[len(input)-1]
	r := []int{piv}

	var s int
	var e int

	//Travel the input array, skipping the last (its already in the correct position)
	for i := 0; i < len(input)-1; i++ {
		v := input[i]

		if v < piv {
			r = append([]int{v}, r...)
			s++
		} else {
			r = append(r, v)
			e++
		}
	}

	return r, s, e
}

func quickSort(input []int) []int {

	a, s, e := partition(input)

	//Before the pivot
	if s > 1 {
		b := quickSort(input[0:s])
		copy(a[0:s+1], b)

	}

	//After the pivot
	if e > 1 {
		q := len(input)
		es := len(input) - e
		c := quickSort(a[es:q])
		copy(a[es:q], c)
	}

	return a

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

	input = quickSort(input)
	fmt.Println("After sort:")
	fmt.Println(input)

	fmt.Println("====")

	return maxPair

}

func StartMaximumProductPair() {
	fmt.Println(maximumProductPair([]int{-10, -3, 7, 6, -2}))
	//	fmt.Println(maximumProductPair([]int{-4, 3, 2, 7, -5}))
	//	fmt.Println(maximumProductPair([]int{1}))

}
