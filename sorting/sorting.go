package main

import (
  "fmt"
  "sort"
)

func main() {

  strs := []string{"c", "a", "d", "b"}

  // sorting strings
  sort.Strings(strs)
  fmt.Println("Sorted:", strs)

  // sorting ints
  ints := []int{1, 3, 2, 4, 6, 5}
  sort.Ints(ints)
  fmt.Println("Sorted ints:", ints)

  // is somthing sorted?
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted?:", s)
}
