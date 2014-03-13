package main

import "fmt"

func main(){
  messages := make(chan string)

  var input string
  fmt.Scanln(&input)
  go func(){messages <- input}()

  msg := <-messages

  fmt.Println(msg)
}
