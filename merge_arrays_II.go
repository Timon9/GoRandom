package main

import "fmt"

/*

Given two sorted integer arrays `X[]` and `Y[]` of size `m` and `n` each where `m >= n` and `X[]` has exactly `n` vacant cells, merge elements of `Y[]` in their correct position in array `X[]`, i.e., merge `(X, Y)` by keeping the sorted order.

Input : Two arrays X[] and Y[] where vacant cells in X[] is represented by 0.

X[] = [0, 2, 0, 3, 0, 5, 6, 0, 0]
Y[] = [1, 8, 9, 10, 15]

Output: X[] = [1, 2, 3, 5, 6, 8, 9, 10, 15]

*/
func solveMergeArraysII(x []int, y []int) {
	for i := 0; i < len(x); i++ {
		v := x[i]
		fmt.Println(v)
	}
}
func MergeArraysII() {
	fmt.Println("MergeArraysII")

	x := []int{0, 2, 0, 3, 0, 5, 6, 0, 0}
	y := []int{1, 8, 9, 10, 15}
	solveMergeArraysII(x, y)
}
