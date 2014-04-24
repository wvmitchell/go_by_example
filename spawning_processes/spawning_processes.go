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

  manCmd := exec.Command("man", "date")
  manStr, err := manCmd.Output()
  if err != nil {
    panic(err)
  }
  p(string(manStr))
}
