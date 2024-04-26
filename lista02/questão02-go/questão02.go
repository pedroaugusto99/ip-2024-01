Questão 02

package main

import (
  "fmt"
)

func main() {
  var popuA, popuB float64
  
  fmt.Print("População País A: ")
  fmt.Scanln(&popuA)
  fmt.Print("População País B: ")
  fmt.Scanln(&popuB)

  anos := 0
  
    for popuA <= popuB {
    popuA = popuA + popuA * 0.03
    popuB = popuB + popuB * 0.015
    anos++
  }

  fmt.Printf("ANOS = %d\n", anos)
}
