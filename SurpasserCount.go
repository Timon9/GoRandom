package main

import "fmt"

/*

Given an integer array having distinct elements, find the surpasser count for each element in it. In other words, for each array element, find the total number of elements to its right, which are greater than it.

Input : [4, 6, 3, 9, 7, 10]
Output: [4, 3, 3, 1, 1, 0]

*/
func surpasserCount(input []int) {
	fmt.Println("Start SurpasserCount")

	hm := make(map[int]int)

	for i := 0; i < len(input); i++ {
		v := input[i]
		hm[v]++
	}
	fmt.Println(hm)

}

func StartSurpasserCount() {
	surpasserCount([]int{4, 6, 3, 9, 7, 10})
}
