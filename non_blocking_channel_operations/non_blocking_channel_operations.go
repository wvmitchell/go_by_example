package main

import "fmt"

func main() {

	messages := make(chan string)
  signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("No message recieved")
	}

  msg := "hi there"
  select {
  case messages <- msg:
    fmt.Println("message sent:", msg)
  default:
    fmt.Println("No message sent")
  }

  select {
  case msg := <-messages:
    fmt.Println("Message received:", msg)
  case sig := <-signals:
    fmt.Println("Signal received:", sig)
  default:
    fmt.Println("No activity")
  }
}
