Questão 08

package main

import (
  "fmt"
)

func main() {
  var x, final int
  fmt.Print("Digite a quantidade de times: ")
  fmt.Scan(&x)

  if x < 2 {
    fmt.Println("Campeonato inválido!")
    return
  }
  
  final = 1
  
  for i := 1; i <= x; i++ {
    for j := i + 1; j <= x; j++ {
      fmt.Printf("Final %d: Time%d X Time%d\n", final, i, j)
      final++
    }
  }
}
