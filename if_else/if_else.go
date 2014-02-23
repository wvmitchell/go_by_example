package main

import "fmt"

func main(){

  if 7%2 != 0 {
    fmt.Println("Odd...")
  } else {
    fmt.Println("Even...")
  }

  if i := 10; i<10 {
    fmt.Println("One digit")
  } else if i>10 {
    fmt.Println("Two digits, bigger than 10")
  } else {
    fmt.Println("Exactly 10")
  }

}
