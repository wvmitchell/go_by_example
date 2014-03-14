package main

import "fmt"
import "time"

func main(){

  timer1 := time.NewTimer(time.Second * 2)

  <-timer1.C
  fmt.Println("Timer 1 expired")


  timer2 := time.NewTimer(time.Second)
}
