package main

import (
	"errors"
	"fmt"
)

/*

Given two sorted integer arrays `X[]` and `Y[]` of size `m` and `n` each where `m >= n` and `X[]` has exactly `n` vacant cells, merge elements of `Y[]` in their correct position in array `X[]`, i.e., merge `(X, Y)` by keeping the sorted order.

Input : Two arrays X[] and Y[] where vacant cells in X[] is represented by 0.

X[] = [0, 2, 0, 3, 0, 5, 6, 0, 0]
Y[] = [1, 8, 9, 10, 15]

Output: X[] = [1, 2, 3, 5, 6, 8, 9, 10, 15]

*/

/*
Function to find the first filled cell from array `x` starting from cursor `c`.
Returns the cursor, value of first cell and error.
*/

func findNextFilledCell(x []int, c int) (int, int, error) {

	for i := c; i < len(x); i++ {
		if x[i] > 0 {
			return c, x[i], nil
		}
		c++
	}
	return 0, 0, errors.New("no next cell found")
}

func solveMergeArraysII(x []int, y []int) {
	m := len(x)
	n := len(y)

	if m < n {
		fmt.Println("Incorrect input array length. Expecting X to be larger than Y")
		return
	}

	/*
		Traverse `x` searching for vacant cells.
	*/

	for i := 0; i < m; i++ {
		v := x[i]

		if v == 0 { // We found a vacant cell

			hi, hn, err := findNextFilledCell(x, i+1) // The highest next cell of X

			if len(y) >= 0 && (y[0] < hn || err != nil) { // The lowest of Y is lower than the next filled cell of X , lets use that.
				x[i] = y[0]

				t := len(y)
				y = y[1:t] // Remove y[0]

			} else if err == nil { // Use the next of X

				x[i] = hn
				x[hi] = v

			}

		}
	}

	fmt.Println(x)
}
func MergeArraysII() {
	fmt.Println("MergeArraysII")

	x := []int{0, 2, 0, 3, 0, 5, 6, 0, 0}
	y := []int{1, 8, 9, 10, 15}
	solveMergeArraysII(x, y)
}
