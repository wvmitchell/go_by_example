package main

import "fmt"
import "math"

type geometery interface {
  area() float64
  perim() float64
}

type square struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (s square) area() float64 {
  return s.width * s.height
}

func (s square) perim() float64 {
  return s.width*2 + s.height*2
}

func (c circle) area() float64 {
  return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

func measure(g geometery) {
  fmt.Println("Shape:", g)
  fmt.Println("Area:", g.area())
  fmt.Println("Perimeter:", g.perim())
}

func main() {
  c := circle{5.0}
  s := square{10, 21}

  measure(c)
  measure(s)
}
