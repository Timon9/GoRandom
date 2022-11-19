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

func find(a []int, k int, s int, m map[mapkey]bool) {

	// fmt.Println("=> Finding", s, "..*")
	// if k == 0 || k > len(a) {
	// 	return
	// }

	j := 0
	r := []int{s}
	for i := 0; i < len(a); i++ {
		v := a[i]
		r = append(r, v)
		//		if _, ok := m[mapkey{s, v}]; !ok {
		// fmt.Println("[", s, ",", v, "]")
		if j == 0 {
			fmt.Println(r)
			r = []int{s}
			j = 0
		} else {
			j--
		}
		// m[mapkey{s, v}] = true
		//		}
	}

}
func solveCombinations(a []int, k int) {

	fmt.Println("Finding combinations k=", k, " ", a)

	if len(a) == 0 && k > len(a) {
		fmt.Println("==invalid==")
		return // Invalid
	}
	if k == 0 {
		return
	}

	m := make(map[mapkey]bool)

	for i := 0; i < len(a); i++ {
		v := a[i]
		b := len(a)
		find(a[i+1:b], k, v, m)

	}
	fmt.Println("=====")

}

func Combinations() {
	fmt.Println("Combinations")

	solveCombinations([]int{2, 3, 4}, 2)
	solveCombinations([]int{1, 2, 1}, 2)
	solveCombinations([]int{1, 1, 1}, 2)
	solveCombinations([]int{1, 2, 3}, 4)
	//solveCombinations([]int{1, 2}, 0)
}

type mapkey struct {
	f int
	s int
}
