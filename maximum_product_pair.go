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

	fmt.Println(input)

	//The pivot point is the last entry in the array
	piv := input[len(input)-1]

	//Build the new array, starting with the pivot point
	r := []int{piv}

	var s, e int

	//Travel the input array, skipping the last (its already in the correct position)
	for i := 0; i < len(input)-1; i++ {
		v := input[i]

		if v < piv { //If its smaller than the pivot, prepend to the new array
			r = append([]int{v}, r...)
			s++ // Count the number of prepends
		} else { //If its bigger than the pivot, append to the new array
			r = append(r, v)
			e++ // Count the number of appends
		}
	}

	//Return the newly build array and the count of prepend and append (before and after pivot)
	return r, s, e
}

func quickSort(input []int) []int {

	a, s, e := partition(input)

	if s > 1 { //Before the first pivot
		b := quickSort(a[0:s]) // Sort before the pivot
		copy(a[0:s+1], b)      // Copy the sorted results into the new ordered array

	}

	if e > 1 { //After the first pivot
		q := len(input)         // The end of the input array
		es := len(input) - e    // The start of the pivot point
		c := quickSort(a[es:q]) // Sort after the first pivot
		copy(a[es:q], c)        // Copy the sorted results into the ordered array
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
	//fmt.Println(maximumProductPair([]int{-10, -3, 7, 6, -2}))
	fmt.Println(maximumProductPair([]int{-4, 3, 2, 7, -5, 4, 6, 3, 2, 2, 4, 6, -1, 3, -4, 5, 6, 7, 8, 4, 1, 3, 5, 3, 7, 8, -1, 5, 4, 3, 1, 4, -4, -1}))
	//	fmt.Println(maximumProductPair([]int{1}))

}
