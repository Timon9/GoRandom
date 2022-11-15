package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No test provided.")
		return
	}

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

	case "mergearrays2":
	case "mergearraysii":
		MergeArraysII()

	case "maximumcontinuoussequence":
		MaximumContinuousSequence()

	case "shufflearrayiii":
		ShuffleArrayIII()

	case "rearrangearray":
		RearrangeArray()

	case "replacearrayelements":
		ReplaceArrayElements()

	case "longestbitonicsubarray":
		LongestBitonicSubarray()

	case "maximumdifferencepair":
		MaximumDifferencePair()

	case "maximumsumsubarray":
		MaximumSumSubarray()

	case "maximumsumsubarrayii":
		MaximumSumSubarrayII()

	case "maximumsumcircularsubarray":
		MaximumSumCircularSubarray()
	}

}
