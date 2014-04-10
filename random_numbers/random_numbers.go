package main

import (
  "fmt"
  "math/rand"
)

func main() {

  p := fmt.Println

  // random ints
  p(rand.Intn(100))
  p(rand.Intn(100))
  p()

  // random floats
  p(rand.Float64())
  p(rand.Float64())
  p()

  // random floats in other ranges
  p(rand.Float64()+5)
  p(rand.Float64()*5+5)
  p()

  // new seed
  s1 := rand.NewSource(50)
  r1 := rand.New(s1)
  p(r1.Intn(100))
  p(r1.Intn(100))
  p()

  // new rand, with same seed, same output
  s2 := rand.NewSource(50)
  r2 := rand.New(s2)
  p(r2.Intn(100))
  p(r2.Intn(100))

}
