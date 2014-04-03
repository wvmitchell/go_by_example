package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
  p("Count:", s.Count("test", "t"))
  p("HasPrefix:", s.HasPrefix("test", "te"))
  p("HasSuffix:", s.HasSuffix("test", "st"))
}
