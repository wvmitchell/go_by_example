package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func main() {

  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)

}
