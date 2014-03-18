package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int){
  for j := range jobs {
    fmt.Prinln("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
}

func main() {
}

