package main

import (
  "fmt"
  "sort"
)

func main() {
  
  var saidas []string
  
  for {
    var N int
    fmt.Scan(&N)
    if N == 0{
      break
    }
    
    vetor := make([]int, N)
    for i := range vetor {
      fmt.Scan(&vetor[i])
    }
    
    sort.Ints(vetor)
    
    frequencia := make([]int, vetor[N-1]+1)
    for _, num := range vetor {
      frequencia[num]++
    }
    
    contagem := 0
    for i, j := range frequencia {
      contagem += j
      saidas = append(saidas, fmt.Sprintf("(%d) %d\n", i, contagem))
    }
  }
  
  for _, saida := range saidas {
    fmt.Print(saida)
  }
}
