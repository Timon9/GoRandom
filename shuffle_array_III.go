package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

Given an array of distinct integer, in-place shuffle the array. The solution should produce an unbiased permutation, i.e., every permutation is equally likely.

Input: [1, 2, 3, 4, 5]
Output: [5, 4, 2, 1, 3] or [4, 1, 5, 3, 2] or any other unbiased permutation.

*/

func solveShuffleArrayIII(numbers []int) []int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < len(numbers); i++ {

		j := i + r.Intn(len(numbers)-i) // Skip the ones we already sorted
		v := numbers[i]

		numbers[i] = numbers[j]
		numbers[j] = v

	}

	return numbers

}
func ShuffleArrayIII() {
	fmt.Println("ShuffleArrayIII")
	fmt.Println(solveShuffleArrayIII([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
