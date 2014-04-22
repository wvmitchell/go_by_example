package main

import (
  "flag"
  "fmt"
)

func main() {

  // name of option, default value, and description
  wordPtr := flag.String("word", "foo", "a string")

  numPtr := flag.Int("num", 45, "an int")

  flag.Parse()

  fmt.Println("wordPtr", *wordPtr)
  fmt.Println("numPtr", *numPtr)
}
