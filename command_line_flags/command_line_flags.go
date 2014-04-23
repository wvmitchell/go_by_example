package main

import (
  "flag"
  "fmt"
)

func main() {

  // name of option, default value, and description
  wordPtr := flag.String("word", "foo", "a string")

  numPtr := flag.Int("num", 45, "an int")

  boolPtr := flag.Bool("fork", false, "a bool")

  flag.Parse()

  // only print if boolPtr has been set to true
  if *boolPtr {
    fmt.Println("wordPtr", *wordPtr)
    fmt.Println("numPtr", *numPtr)
    fmt.Println("fork", *boolPtr)
  }
}
