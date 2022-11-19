package main

import "fmt"

/*

Given an integer array, find all distinct combinations of a given length `k`.
The solution should return a set containing all the distinct combinations, while preserving the relative order of elements as they appear in the array.

Input : [2, 3, 4], k = 2
Output: {[2, 3], [2, 4], [3, 4]}

Input : [1, 2, 1], k = 2
Output: {[1, 2], [1, 1], [2, 1]}

Input : [1, 1, 1], k = 2
Output: {[1, 1]}

Input : [1, 2, 3], k = 4
Output: {}

Input : [1, 2], k = 0
Output: {[]}

*/

func find(a []int, k int, s int) {
	for i := 0; i < len(a); i++ {
		v := a[i]
		fmt.Println("[", s, ",", v, "]")
	}

}
func solveCombinations(a []int, k int) {

	if len(a) == 0 && k > len(a) {
		return // Invalid
	}
	if k == 0 {
		return
	}

	for i := 0; i < len(a); i++ {
		v := a[i]
		b := len(a)
		find(a[i+1:b], k, v)

	}
	fmt.Println("=====")

}

func Combinations() {
	fmt.Println("Combinations")

	solveCombinations([]int{2, 3, 4}, 2)
	solveCombinations([]int{1, 2, 1}, 2)
	solveCombinations([]int{1, 1, 1}, 2)
	//solveCombinations([]int{1, 2, 3}, 4)
	//solveCombinations([]int{1, 2}, 0)
}
