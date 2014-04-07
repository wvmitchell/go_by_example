package main

import (
  "regexp"
  "fmt"
  "bytes"
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

  // All variants
  pf(re.FindAllString("peach punch patch pinch", -1))
  pf(re.FindAllStringSubmatch("peach pinch punch", -1))
  pf(re.FindAllStringSubmatchIndex("peach punch pinch", -1))
  pf(re.FindAllString("peach punch patch pinch", 2))


  // byte variants
  pf(re.Match([]byte("peach")))
  pf(re.Find([]byte("peach punch patch pinch")))

  // MustCompile for setting constants, only one return value, unlike Compile
  re = regexp.MustCompile("p([a-z]+)ch")
  pf(re)

  pf(re.ReplaceAllString("A pretty peach", "pear"))

  // ReplaceAllFunc allows transformation of matched text with function
  in := []byte("a peach")
  out := re.ReplaceAllFunc(in, bytes.ToUpper)
  pf(string(out))

}
