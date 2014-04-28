package main

import (
  "fmt"
  "os"
)

func main() {

  // defer is not run when using exit, so ! will never be printed
  defer fmt.Println("!")

  os.Exit(3)
}
