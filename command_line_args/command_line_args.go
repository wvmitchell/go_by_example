package main

import (
  "os"
  "fmt"
)

func main() {

  // aruguments with name of program and path
  argsWithProg := os.Args

  // and without
  argsWithoutProg := os.Args[1:]

  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)

}
