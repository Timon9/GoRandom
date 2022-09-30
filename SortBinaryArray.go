package main

import "fmt"

/*
Status: Done in O(n)

Given a binary array, in-place sort it in linear time and constant space. The output should contain all zeroes, followed by all ones.

Input : [1, 0, 1, 0, 1, 0, 0, 1]
Output: [0, 0, 0, 0, 1, 1, 1, 1]

Input : [1, 1]
Output: [1, 1]


*/

func prependInt(x []int, y int) []int {
	//Create new array +1 larger than x
	a := make([]int, len(x)+1)
	//Set the first key to value y
	a[0] = y
	//Copy the original array to a, skipping key 0
	copy(a[1:], x)
	return a

}

func sortBinaryArray(nums []int) {
	fmt.Println("SortBinaryArray")

	var r []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			r = append(r, 1)
		} else {
			//r = append([]int{0}, r...)
			r = prependInt(r, 0)
		}

	}
	fmt.Println(r)

}

func StartSortBinaryArray() {
	sortBinaryArray([]int{1, 0, 1, 0, 1, 0, 0, 1})
}
