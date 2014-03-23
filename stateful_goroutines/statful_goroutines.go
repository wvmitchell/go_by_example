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
}
