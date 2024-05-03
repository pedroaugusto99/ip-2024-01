package main

import "fmt"

func main() {
  
  var n int
  fmt.Println("Entrada")
  fmt.Scanln(&n)
  
  numeros := make([]int, n)
  for i := 0;i < n; i++{
    fmt.Scan(&numeros[i])
  } 
  
  frequencia := make(map[int]int)
  for _, numero := range numeros {
    frequencia[numero]++
  }

 
  unico := 0
  for _, elemento := range frequencia {
    if elemento == 1 {
      unico++
    }
  }
  fmt.Println(Saída)
  fmt.Println(unico)
}
