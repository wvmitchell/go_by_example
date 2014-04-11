package main

import (
  "strconv"
  "fmt"
)

func main() {

  p := fmt.Println

  // parse float
  f, _ := strconv.ParseFloat("2.323", 64)
  p(f)

  // parse int
  i, _ := strconv.ParseInt("234a", 12, 64)
  p(i)

}
