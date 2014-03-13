package main

import "fmt"
import "time"

func worker(done chan bool){
  fmt.Println("working...")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true
}

func counter(count chan bool){
  for i := 0; i< 5; i++ {
    fmt.Println(i)
  }
  count <-true
}

func main(){

  done := make(chan bool)
  go worker(done)
  go counter(done)

  <-done
}
