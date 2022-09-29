package main

import "fmt"

/*

Given an integer array having distinct elements, find the surpasser count for each element in it. In other words, for each array element, find the total number of elements to its right, which are greater than it.

Input : [4, 6, 3, 9, 7, 10]
Output: [4, 3, 3, 1, 1, 0]

https://www.techiedelight.com/surpasser-count-each-element-array/

*/

func surpasserCount(input []int) {
	fmt.Println("Start SurpasserCount")

	hm := make(map[int]int)

	l := len(input)

	for i := 0; i < l; i++ {
		v := input[i]

		for ii := 1; ii < v; ii++ {
			hm[ii]++
		}

	}

	fmt.Println(hm)

}

func StartSurpasserCount() {
	surpasserCount([]int{4, 6, 3, 9, 7, 10})
}
