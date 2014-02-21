package main

import "fmt"

func main(){

  i := 0
  for i <= 3 {
    fmt.Println(i)
    i++
  }

  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      fmt.Println(i)
    } else {
      fmt.Println("it's odd brah")
    }
  }
}
