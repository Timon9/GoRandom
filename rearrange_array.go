package main

import "fmt"

/*

Given an integer array, rearrange it such that every second element becomes greater than its left and right elements.

• Assume that no duplicate elements are present in the input array.
• The solution should perform single traveral of the array.
• In case the multiple rearrangement exists, the solution can return any one of them.

Input : [1, 2, 3, 4, 5, 6, 7]
Output: [1, 3, 2, 5, 4, 7, 6] or [1, 5, 2, 6, 3, 7, 4], or any other valid combination..

Input : [6, 9, 2, 5, 1, 4]
Output: [6, 9, 2, 5, 1, 4] or [1, 5, 2, 6, 4, 9], or any other valid combination..

*/
func solveRearrangeArray(numbers []int) {

	for i := 1; i < len(numbers)-1; i = i + 2 {

		v := numbers[i]

		l := numbers[i-1]
		r := numbers[i+1]

		if v < l {
			// Swap l and v
			tmp := numbers[i-1]
			numbers[i-1] = numbers[i]
			numbers[i] = tmp
		}

		if v < r {
			// Swap r and v
			numbers[i+1] = v
			numbers[i] = r
		}
	}

	fmt.Println(numbers)
}

func RearrangeArray() {
	fmt.Println("RearrangeArray")
	solveRearrangeArray([]int{1, 2, 3, 4, 5, 6, 7})
}
