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

// Function to return true if all in collection satisfy predicate f
func All(vs []string, f func(string) bool) bool {
  for _, v := range vs {
    if !f(v) {
      return false
    }
  }
  return true
}

// Function returns new slice of all elements which satisfy precicate f
func Filter(vs []string, f func(string) bool) (rs []string) {
  for _, v := range vs {
    if f(v) {
      rs = append(rs, v)
    }
  }
  return rs
}

// Function returns new slice of elements which result after applying function f
func Collect(vs []string, f func(string) string) (rs []string) {
  for _, v := range vs {
    rs = append(rs, f(v))
  }
  return rs
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

  fmt.Println(Filter(words, func(v string) bool {
    return strings.Contains(v, "e")
  }))

  fmt.Println(Collect(words, func(v string) string {
    return strings.ToUpper(v)
  }))
}
