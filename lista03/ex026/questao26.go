package main

import (
  "fmt"
  "sort"
)

func anoes(numeros []int) []int {
  sort.Ints(numeros)
  var resultado []int
  for _, num := range numeros {
    if num <= 100 {
      resultado = append(resultado, num)
      if len(resultado) == 7 {
        break
      }
    }
  }
  return resultado
}

func main() {
  var T int
  fmt.Scan(&T)
  for t := 0; t < T; t++ {
    var numeros [9]int
    for i := 0; i < 9; i++ {
      fmt.Scan(&numeros[i])
    }
    resultado := anoes(numeros[:])
    for _, num := range resultado {
      fmt.Println(num)
    }
  }
}
