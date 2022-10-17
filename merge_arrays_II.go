package main

import "fmt"

/*

Given two sorted integer arrays `X[]` and `Y[]` of size `m` and `n` each where `m >= n` and `X[]` has exactly `n` vacant cells, merge elements of `Y[]` in their correct position in array `X[]`, i.e., merge `(X, Y)` by keeping the sorted order.

Input : Two arrays X[] and Y[] where vacant cells in X[] is represented by 0.

X[] = [0, 2, 0, 3, 0, 5, 6, 0, 0]
Y[] = [1, 8, 9, 10, 15]

Output: X[] = [1, 2, 3, 5, 6, 8, 9, 10, 15]

*/

func findNextFilledCell(x []int) (int, error) {
	for i := 0; i < len(x); i++ {
		if x[i] > 0 {
			return x[i], nil
		}
	}
	return 0, nil
}

func solveMergeArraysII(x []int, y []int) {
	m := len(x)
	n := len(y)

	if m < n {
		fmt.Println("Incorrect input array length. Expecting X to be larger than Y")
		return
	}

	for i := 0; i < m; i++ {
		v := x[i]
		r := y[0]

		if v == 0 { // We found a vacant cell
			if (i + 1) < m { // We have a next cell

				hn, err := findNextFilledCell(x[i:m]) // The highest next cell of X

				if len(y) > 0 && (r < hn || err == nil) { // The lowest of Y is lower than the next filled cell of X , lets use that.
					fmt.Printf("Swapping from Y #%d for %d\n", i, r)
				} else if err == nil { // Use the next of X
					fmt.Printf("Swapping from X #%d for %d\n", i, hn)
				}

			}
		}
	}
}
func MergeArraysII() {
	fmt.Println("MergeArraysII")

	x := []int{0, 2, 0, 3, 0, 5, 6, 0, 0}
	y := []int{1, 8, 9, 10, 15}
	solveMergeArraysII(x, y)
}
