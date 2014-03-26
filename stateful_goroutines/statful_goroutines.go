package main

import (
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  val int
  resp chan bool
}

func main() {

  var ops int64 = 0;

  reads := make(chan *readOp)
  writes := make(chan *writeOp)

  go func() {
    var state = make(map[int][int])

    for {
      select {
        case read := <-reads: 
          read.resp <- state[read.key]
        case writes := <-writes:
          state[write.key] = write.val
          write.resp <- true
        }
      }
    }
  }
}
