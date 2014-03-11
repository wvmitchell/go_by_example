package main

import "fmt"

type rect struct {
  width, height int
}

func (r *rect) area() int {
  return r.width * r.height
}

func main(){

  r := rect{10, 10}
  fmt.Println("rectangle area:", r.area())

}

