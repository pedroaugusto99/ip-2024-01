package main

import (
  "fmt"
  "sort"
)

func main() {
  
  var saidas []string
  
  var T int
  fmt.Scan(&T)

  for i := 0; i < T; i++ {
    var N int
    fmt.Scan(&N)

    numeros := make([]int, N)
    for i := 0; i < N; i++ {
      fmt.Scan(&numeros[i])
    }

    sort.Ints(numeros)

    distMin := numeros[1] - numeros[0]
    for i := 2; i < N; i++ {
      distancia := numeros[i] - numeros[i-1]
      if distancia < distMin {
        distMin = distancia
      }
    }

    comparacoes := (N - 1) * N / 2
    
    saidas = append(saidas,fmt.Sprintf("%d %d\n", distMin, comparacoes))
  }
  for _, saida := range saidas {
    fmt.Print(saida)
}
}
