package main

import (
  "fmt"
  "os"
)

func main() {

  f := createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)

}

func createFile(p string) *os.File {
  fmt.Println("Creating")
  f, err := os.Create(p)
  if err != nil {
    panic(err)
  }
  return f
}
