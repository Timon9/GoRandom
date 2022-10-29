package main

import "fmt"

/*

Status: Done in O(n)

Given an integer array of size `n`, return the element which appears more than `n/2` times. Assume that the input always contain the majority element.

Input : [2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2]
Output: 2

Input : [1, 3, 1, 1]
Output: 1

*/

func findMajorityElement(input []int) int {
	fmt.Printf("\nStartFindMajorityElement\n====\nInput:%d\n===\n\n", input)

	hm := make(map[int]int)
	a := len(input) / 2

	for i := 0; i < len(input); i++ {
		v := input[i]
		hm[v]++
		if hm[v] > a {
			return v
		}
	}
	return 0
}
func findMajorityElementBoyerMooreAlg(input []int) int { // No hashmap but boyer-moore alg
	fmt.Printf("\nStartFindMajorityElement\n====\nInput:%d\n===\n\n", input)

	m := 0 // Majority Element
	c := 0 // Counter

	for i := 0; i < len(input); i++ {
		v := input[i]
		if c == 0 {
			m = v
			c++
		} else if m == v {
			c++
		} else {
			c--
		}

	}
	return m
}
func StartFindMajorityElement() {
	t1 := findMajorityElement([]int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2})
	fmt.Println(t1)

	t2 := findMajorityElementBoyerMooreAlg([]int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2})
	fmt.Println(t2)
}
