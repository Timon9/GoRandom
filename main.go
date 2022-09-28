package main

import (
	"fmt"
)

/*
Given an integer array, find a subarray having a given sum in it.

For example,

Input:  nums[] = {2, 6, 0, 9, 7, 3, 1, 4, 1, 10}, target = 15
Output: {6, 0, 9}


Input:  nums[] = {0, 5, -7, 1, -4, 7, 6, 1, 4, 1, 10}, target = 15
Output: {1, -4, 7, 6, 1, 4} or {4, 1, 10}


Input:  nums[] = {0, 5, -7, 1, -4, 7, 6, 1, 4, 1, 10}, target = -3
Output: {1, -4} or {-7, 1, -4, 7}
*/

func findSubArrayContainingSum(arr []int, sum int) []int {
	fmt.Printf("Searching for sum %d subarrays.\n", sum)

	r := []int{}

	for i := 0; i < len(arr); i++ {
		v := arr[i]
		r = append(r, v)
	}
	return r
}

func isSame(r []int, e []int) bool {
	if len(r) != len(e) {
		return false
	}
	for i := 0; i < len(e); i++ {
		if r[i] != e[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello again!")

	arr := []int{2, 6, 0, 9, 7, 3, 1, 4, 1, 10}
	r := findSubArrayContainingSum(arr, 15)
	e := []int{6, 0, 9}

	if isSame(r, e) {
		fmt.Println("Succes!")
	} else {
		fmt.Printf("Failed. Expected %d, got %d.\n", e, r)
	}

}
