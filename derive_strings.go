package main

import "fmt"

/*

Check if a given string can be derived from another string by circularly rotating it. The rotation can be in a clockwise or anti-clockwise rotation.

Input: X = "ABCD", Y = "DABC"
Output: true
Explanation: "DABC" can be derived from "ABCD" by right-rotating it by 1 unit

*/

func deriveStrings(haystack string, needle string) {
	fmt.Printf("DeriveStrings: Searching for %v in %v\n\n", needle, haystack)
	for i := 0; i < len(haystack); i++ {
		v := haystack[i]
		if v == needle[0] {
			c := i
			for ii := 1; ii < len(needle); ii++ {
				c = i + ii
				if c > len(haystack)-1 {
					c = c % len(haystack)
				}

				if needle[ii] != haystack[c] {
					break
				}
				if ii == len(needle)-1 {
					fmt.Printf("Found match! Started at %d ended at %d\n\n", i, c)
				}

			}

		}

	}
}

func DeriveStrings() {
	deriveStrings("ABCD", "DABCDABCDABCDABCDA")
}
