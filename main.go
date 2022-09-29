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
}
