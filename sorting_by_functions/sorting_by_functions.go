package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	fruits := []string{"apple", "pear", "watermellon", "peach", "orange", "bananna"}

	fmt.Println("Unsorted:", fruits)
	sort.Sort(ByLength(fruits))
	fmt.Println("Sorted:", fruits)

}
