package main

import "fmt"

/*

Given an integer array of size `n`, return the element which appears more than `n/2` times. Assume that the input always contain the majority element.

Input : [2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2]
Output: 2

Input : [1, 3, 1, 1]
Output: 1

*/

func findMajorityElement(input []int) {
	fmt.Printf("StartFindMajorityElement\n====\nInput:%d\n===\n\n", input)

}
func StartFindMajorityElement() {
	findMajorityElement([]int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2})
	findMajorityElement([]int{1, 3, 1, 1})
}
