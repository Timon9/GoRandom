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

func findSubArrayContainingSum(arr []int, sum int) map[int][]int {
	fmt.Printf("Searching for sum %d subArrays in %d.\n===\n", sum, arr)

	m := make(map[int]int)
	r := make(map[int][]int)

	m[0] = -1
	var count_so_far int

	for i := 0; i < len(arr); i++ {
		v := arr[i]
		count_so_far += v
		n := count_so_far - sum
		if _, ok := m[n]; ok {
			rs := arr[m[n]+1 : i+1]
			r[len(r)+1] = rs

			fmt.Printf("Found subArray: %d\n", rs)
		}
		m[count_so_far] = i
	}
	return r
}

/*
func isSame(r map[int][]int, e map[int][]int) bool {

	//Check if the expected map has the same size as the result map.
	if len(r) != len(e) {
		return false
	}

	//Iterate the map
	for i := 0; i < len(e); i++ {
		//Loop the result
		for ii := 0; ii < len(e[i]); ii++ {
			if r[i][ii] != e[i][ii] {
				return false
			}

		}
	}

	//No issues found
	return true
}
*/

func StartFindSubArrayContainingSum() {

	arr := []int{2, 6, 0, 9, 7, 3, 1, 4, 1, 10}
	findSubArrayContainingSum(arr, 15)

}
