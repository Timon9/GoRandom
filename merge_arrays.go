package main

import "fmt"

/*

Given two sorted integer arrays, `X[]` and `Y[]` of size `m` and `n` each, in-place merge elements of `X[]` with elements of array `Y[]` by maintaining the sorted order, i.e., fill `X[]` with the first `m` smallest elements and fill `Y[]` with remaining elements.

Input :

X[] = [1, 4, 7, 8, 10]
Y[] = [2, 3, 9]

//

X[] = [1, 4, 7, 8, 10]
Y[] = [2, 3, 9]


Output:

X[] = [1, 2, 3, 4, 7]
Y[] = [8, 9, 10]

*/

func solve(x []int, y []int) ([]int, []int) {

	for i := 0; i < len(x); i++ {

		if y[0] < x[i] {
			// Swap the results of lowest y with current x
			tmp := x[i]
			x[i] = y[0]
			y[0] = tmp

			// Put the swapped y[0] in the correct position

			for ii := 1; ii < len(y); ii++ {
				if y[ii] < y[ii-1] { // Next number is lower? swap
					tmp := y[ii-1]
					y[ii-1] = y[ii]
					y[ii] = tmp
				}
			}

		}
	}

	fmt.Println(x)
	fmt.Println(y)

	return x, y
}

func MergeArrays() {
	fmt.Println("MergeArrays")
	solve([]int{1, 4, 7, 8, 10}, []int{2, 3, 9})
}
