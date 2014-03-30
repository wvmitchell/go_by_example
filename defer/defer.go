package main

import (
  "fmt"
  "os"
)

func main() {

  f := createFile("/tmp/defer.txt")

  // closeFile will be defered until the end of execution
  // defer statments are first in last out
  defer endOfProgram()
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

func writeFile(f *os.File) {
  fmt.Println("Writing")
  fmt.Fprintln(f, "here is some data")
}

func closeFile(f *os.File) {
  fmt.Println("Closing")
  f.Close()
}

func endOfProgram() {
  fmt.Println("This is the end")
}
