package main

import (
	"os"
	"strings"
)

func main() {

	runTest := strings.ToLower(os.Args[1])

	if runTest == "findsubarraycontainingsum" {
		StartFindSubArrayContainingSum()
	}
	if runTest == "startfindmajorityelement" {
		StartFindMajorityElement()
	}
	if runTest == "repeatedsubsequence" {
		StartRepeatedSubsequence()
	}
	if runTest == "startsortbinaryarray" {
		StartSortBinaryArray()
	}

}
