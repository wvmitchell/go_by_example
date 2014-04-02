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

// Function to return true if any in collection satisfy predicate f
func Any(vs []string, f func(string) bool) bool {
  for _, v := range vs {
    if f(v) {
      return true
    }
  }
  return false
}

func All(vs []string, f func(string) bool) bool {
  for _, v := range vs {
    if !f(v) {
      return false
    }
  }
  return true
}

func main() {
  words := strings.Split("the quick brown fox jumps over the lazy dog", " ")
  fmt.Println(Index(words, "fox"))
  fmt.Println(Include(words, "boat"))
  fmt.Println(Include(words, "brown"))

  fmt.Println(Any(words, func(v string) bool {
    return len(v) == 40
  }))

  fmt.Println(All(words, func(v string) bool {
    return strings.ToLower(v) == v
  }))

}
