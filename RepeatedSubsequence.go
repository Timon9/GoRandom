package main

import "fmt"

/*

Given a string, check if a repeated subsequence is present in it or not. The repeated subsequence should have a length of 2 or more.

Input: "XYBAXB"
Output: true
Explanation: String "XYBAXB" has XB(XBXB) as a repeated subsequence

Input: "XBXAXB"
Output: true
Explanation: String "XBXAXB" has XX(XXX) as a repeated subsequence

Input: "XYBYAXBY"
Output: true
Explanation: String "XYBYAXBY" has XB(XBXB), XY(XYXY), YY(YYY), YB(YBYB), and YBY(YBYBY) as repeated subsequences.

Input: "ABCA"
Output: false
Explanation: String "ABCA" doesn’t have any repeated subsequence

*/
func repeatedSubsequence(input string) {

	hm := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		v := input[i]
		hm[v]++
		fmt.Printf("%c [%d]\n", v, v)
	}
	fmt.Println("===")
	fmt.Println(hm)
}

func StartRepeatedSubsequence() {
	fmt.Println("Start RepeatedSubsequence")

	repeatedSubsequence("XYBAXB")
}
