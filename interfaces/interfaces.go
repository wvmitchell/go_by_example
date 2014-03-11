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
