package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func main() {

  p := fmt.Println

  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)

  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    sig := <-sigs
    p()
    p(sig)
    done <- true
  }()

  p("Awaiting signal")
  <-done
  p("Exiting")

}
