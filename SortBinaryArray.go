package main

import "fmt"

/*
Given a binary array, in-place sort it in linear time and constant space. The output should contain all zeroes, followed by all ones.

Input : [1, 0, 1, 0, 1, 0, 0, 1]
Output: [0, 0, 0, 0, 1, 1, 1, 1]

Input : [1, 1]
Output: [1, 1]
*/
func sortBinaryArray(nums []int) {
	fmt.Println("SortBinaryArray")

	for i := 0; i <= len(nums); i++ {
		fmt.Printf("%d\n", i)
	}
}

func StartSortBinaryArray() {
	sortBinaryArray([]int{1, 0, 1, 0, 1, 0, 0, 1})
}
