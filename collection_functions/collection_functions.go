package main

import (
	"fmt"
	"strings"
)

// Function to return first index of the target string

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func main() {
}
