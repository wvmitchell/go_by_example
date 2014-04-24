package main

import (
  "fmt"
  //"io/ioutil"
  "os/exec"
)

func main() {

  p := fmt.Println

  dateCmd := exec.Command("date")

  dateOut, err := dateCmd.Output()
  if err != nil {
    panic(err)
  }
  p("> date")
  p(string(dateOut))
}
