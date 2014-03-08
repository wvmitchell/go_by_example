package main

import "fmt"

func main(){

  m := make(map[string]int)

  m["one"] = 1
  m["two"] = 2

  fmt.Println(m)

  v1 := m["one"]
  fmt.Println("v1: ", v1)

  delete(m, "one")
  fmt.Println(m)

  val, prs := m["three"]

  fmt.Println(val, prs)

  n := map[string]int{"one":1, "two":2}
  fmt.Println("map: ", n)
}
