package main

import "fmt"

func main() {
  var n int
  fmt.Println("Entrada")
  fmt.Scanln(&n)
  array := make([]int, n, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&array[i])
  }
  fmt.Print("Saída\n") 
  soma := 0
  for _, numeros := range array {
    soma += numeros
  }
  fmt.Println(soma)
}
