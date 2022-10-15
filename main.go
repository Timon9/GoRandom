package main

import (
	"os"
	"strings"
)

func main() {

	runTest := strings.ToLower(os.Args[1])

	switch runTest {

	case "findsubarraycontainingsum":
		StartFindSubArrayContainingSum()

	case "startfindmajorityelement":
		StartFindMajorityElement()

	case "repeatedsubsequence":
		StartRepeatedSubsequence()

	case "startsortbinaryarray":
		StartSortBinaryArray()

	case "maximumproductpair":
		StartMaximumProductPair()

	case "derivestrings":
		DeriveStrings()

	case "dutchnationalflagproblem":
		DutchNationalFlagProblem()

	case "mergearrays":
		MergeArrays()

	}

}
