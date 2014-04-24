package main

import (
  "fmt"
  "io/ioutil"
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

  // won't currently run, for brevity
  if false {
    manCmd := exec.Command("man", "date")
    manStr, err := manCmd.Output()
    if err != nil {
      panic(err)
    }
    p(string(manStr))
  }

  grepCmd := exec.Command("grep", "hello")

  grepIn, _ := grepCmd.StdinPipe()
  grepOut, _ := grepCmd.StdoutPipe()
  grepCmd.Start()
  grepIn.Write([]byte("hello grep\ngoodbye grep"))
  grepIn.Close()
  grepBytes, _ := ioutil.ReadAll(grepOut)
  grepCmd.Wait()

  p("> grep")
  p(string(grepBytes))
}
