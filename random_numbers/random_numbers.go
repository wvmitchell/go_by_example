package main

import (
  "fmt"
  "math/rand"
)

func main() {

  p := fmt.Println

  p(rand.Intn(100))
  p(rand.Intn(100))

  p(rand.Float64())
  p(rand.Float64())

}
