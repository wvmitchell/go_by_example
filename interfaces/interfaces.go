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

func (s *square) area() float64 {
  return s.width * s.height
}

func (s *square) perim() float64 {
  return s.width*2 + s.height*2
}
