package main

import (
  "regexp"
  "fmt"
)

var pf = fmt.Println

func main() {

  // returns true if string matches regexp
  match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
  fmt.Println(match)

  // more often, we'll compile a regexp for reuse
  re, _ := regexp.Compile("p([a-z]+)ch")

  // match test like earlier
  pf(re.MatchString("peach"))

  // find the first match
  pf(re.FindString("peach punch"))

  // also finds first match, but returns start and end indexes of string
  pf(re.FindStringIndex("peach punch"))

  // finds info for first match and first submatch
  pf(re.FindStringSubmatch("peach punch"))

  // finds indexes for first match and first submatch
  pf(re.FindStringSubmatchIndex("peach punch"))


}
