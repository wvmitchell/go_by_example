package main

import "fmt"

func main(){

  nums := []int{1,2,3}

  for _, num := range nums{
    fmt.Println(num)
  }

  m := make(map[string]int)
  m["red"] = 1
  m["blue"] = 2
  for k, v := range m{
    fmt.Println("key:", k, "value:", v)
  }

  letters := "letters"

  // Unicode code points
  for i, c := range letters {
    fmt.Println(i, c)
  }
}
