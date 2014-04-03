package main

import (
	"fmt"
  // alias strings package
	s "strings"
)

// alias fmt.Println method
var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
  p("Count:", s.Count("test", "t"))
  p("HasPrefix:", s.HasPrefix("test", "te"))
  p("HasSuffix:", s.HasSuffix("test", "st"))
  p("Index:", s.Index("test", "e"))
  p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
  p("Repeat:", s.Repeat("ab", 5))
  p("Replace:", s.Replace("foofooberry", "oo", "ool", 1))
  p("Replace:", s.Replace("foofooberry", "oo", "ool", -1))
  p("Split:", s.Split("one fine day", " "))
  p("ToUpper:", s.ToUpper("once upon a time"))
  p("ToLower:", s.ToLower("THERE WAS A RABBIT IN A TRAP"))

  // Theses are not part of the strings package, but are useful when working with strings
  p("Len:", len("a string"))
  p("Char:", "string"[1])
  p("Letter:", string("string"[1]))
}

