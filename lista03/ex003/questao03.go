package main

import "fmt"

func main() {
  var n int
  fmt.Println("Entrada")
  fmt.Scan(&n)

  numeros := make([]int,n,n)
  for i := range numeros {
    fmt.Scan(&numeros[i])
  }

  fmt.Println("Saída")
  for i := n - 1; i >= 0; i-- {
    fmt.Printf("%v ",numeros[i])
  }
}
