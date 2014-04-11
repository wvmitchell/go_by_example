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

  // parse int (base value of 0 requires inference of base from string)
  i, _ := strconv.ParseInt("234", 0, 64)
  p(i)

  n, _ := strconv.ParseInt("-234", 0, 64)
  p(n)

  // parse hex
  h, _ := strconv.ParseInt("0x2a3c4f", 0, 64)
  p(h)

  // parse uint
  u, _ := strconv.ParseUint("483", 0, 64)
  p(u)

  // Atoi convinience method for base 10 numbers
  t, _ := strconv.Atoi("43")
  p(t)

  // Atoi throws error on invalid input
  _, e := strconv.Atoi("potato")
  p(e)

}
