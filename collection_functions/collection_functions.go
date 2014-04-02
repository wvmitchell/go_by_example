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

// Function to return true if the target string is in the slice of strings
func Include(vs []string, t string) bool {
  for _, v := range vs {
    if v == t {
      return true
    }
  }
  return false
}

func main() {
  words := strings.Split("The quick brown fox jumps over the lazy dog", " ")
  fmt.Println(Index(words, "fox"))
  fmt.Println(Include(words, "boat"))
  fmt.Println(Include(words, "brown"))
}
